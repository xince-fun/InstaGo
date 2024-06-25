package app

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/user/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/user/domain/repo"
	"github.com/xince-fun/InstaGo/server/services/user/infra/cache"
	"github.com/xince-fun/InstaGo/server/services/user/infra/sal"
	"github.com/xince-fun/InstaGo/server/services/user/pkg/md5"
	"github.com/xince-fun/InstaGo/server/services/user/pkg/paseto"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/blob"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"time"

	hpaseto "github.com/hertz-contrib/paseto"
)

var UserApplicationServiceSet = wire.NewSet(
	repo.UserRepositorySet,
	md5.EncryptManagerSet,
	paseto.TokenGeneratorSet,
	sal.BlobManagerSet,
	cache.CacheManagerSet,
	NewUserApplicationService,
	wire.Bind(new(EncryptManager), new(*md5.EncryptManager)),
	wire.Bind(new(TokenGenerator), new(*paseto.TokenGenerator)),
	wire.Bind(new(BlobManager), new(*sal.BlobManager)),
	wire.Bind(new(CacheManager), new(*cache.RedisManager)),
)

type UserApplicationService struct {
	userRepo       repo.UserRepository
	encryptManager EncryptManager
	tokenGenerator TokenGenerator
	blobManager    BlobManager
	cacheManager   CacheManager
}

type EncryptManager interface {
	EncryptPassword(string) string
}

type TokenGenerator interface {
	CreateToken(*hpaseto.StandardClaims) (string, error)
}

type BlobManager interface {
	UploadBlob(context.Context, *blob.GeneratePutPreSignedUrlRequest) (*blob.GeneratePutPreSignedUrlResponse, error)
	GetBlob(context.Context, *blob.GenerateGetPreSignedUrlRequest) (*blob.GenerateGetPreSignedUrlResponse, error)
	NotifyBlobUpload(context.Context, *blob.NotifyBlobUploadRequest) (*blob.NotifyBlobUploadResponse, error)
}

type CacheManager interface {
	Get(context.Context, string, interface{}) error
	Set(context.Context, string, cache.CacheItem) error
}

func NewUserApplicationService(userRepo repo.UserRepository, encryptManager EncryptManager,
	tokenGenerator TokenGenerator, blobManager BlobManager, cacheManager CacheManager) *UserApplicationService {
	return &UserApplicationService{
		userRepo:       userRepo,
		encryptManager: encryptManager,
		tokenGenerator: tokenGenerator,
		blobManager:    blobManager,
		cacheManager:   cacheManager,
	}
}

func (s *UserApplicationService) RegisterPhone(ctx context.Context, req *user.RegisterPhoneRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	if r, _ := s.userRepo.FindUserAccountByPhoneNumberNonNil(ctx, req.PhoneNumber); r != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RecordExist)
		return resp, nil
	}

	usrID := uuid.UUID{}
	if usrID, err = s.userRepo.NextIdentity(); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	usr := entity.NewUser(usrID)

	if err = usr.UserAccount.SetPhoneNumber(req.PhoneNumber); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	usr.UserAccount.SetPasswd(s.encryptManager.EncryptPassword(req.Passwd))
	usr.UserInfo.SetFullName(req.FullName)
	usr.UserInfo.SetAccount(req.Account)

	token, err := s.registerUser(ctx, usr)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError.WithMessage(err.Error()))
		return resp, nil
	}

	resp.Token = token
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) RegisterEmail(ctx context.Context, req *user.RegisterEmailRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	if r, _ := s.userRepo.FindUserAccountByEmailNonNil(ctx, req.Email); r != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.RecordExist)
		return resp, nil
	}

	userID := uuid.UUID{}
	if userID, err = s.userRepo.NextIdentity(); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	usr := entity.NewUser(userID)

	if err = usr.UserAccount.SetEmail(req.Email); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}
	usr.UserAccount.SetPasswd(s.encryptManager.EncryptPassword(req.Passwd))
	usr.UserInfo.SetFullName(req.FullName)
	usr.UserInfo.SetAccount(req.Account)

	token, err := s.registerUser(ctx, usr)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError.WithMessage(err.Error()))
		return resp, nil
	}

	resp.Token = token
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) LoginPhone(ctx context.Context, req *user.LoginPhoneRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	usrAccount, err := s.userRepo.FindUserAccountByPhoneNumberNonNil(ctx, req.Account)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	if s.encryptManager.EncryptPassword(req.Passwd) != usrAccount.Passwd {
		resp.BaseResp = utils.BuildBaseResp(errno.UserPwdError)
		return resp, nil
	}

	token, err := s.generateToken(usrAccount.UserID.String())
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	resp.Token = token
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) LoginEmail(ctx context.Context, req *user.LoginEmailRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	usrAccount, err := s.userRepo.FindUserAccountByEmailNonNil(ctx, req.Account)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	if s.encryptManager.EncryptPassword(req.Passwd) != usrAccount.Passwd {
		resp.BaseResp = utils.BuildBaseResp(errno.UserPwdError)
		return resp, nil
	}

	token, err := s.generateToken(usrAccount.UserID.String())
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	resp.Token = token
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) registerUser(ctx context.Context, usr *entity.User) (token string, err error) {
	if err = s.userRepo.SaveUser(ctx, usr); err != nil {
		return "", err
	}

	return s.generateToken(usr.UserID.String())
}

func (s *UserApplicationService) generateToken(id string) (token string, err error) {
	now := time.Now()
	token, err = s.tokenGenerator.CreateToken(&hpaseto.StandardClaims{
		ID:        id,
		Issuer:    consts.Issuer,
		Audience:  consts.User,
		IssuedAt:  now,
		NotBefore: now,
		ExpiredAt: now.Add(consts.FiftyDays),
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserApplicationService) UpdateEmail(ctx context.Context, req *user.UpdateEmailRequest) (resp *user.UpdateEmailResponse, err error) {
	resp = new(user.UpdateEmailResponse)

	usrAccount, err := s.userRepo.FindUserAccountByUserIDNonNil(ctx, req.UserId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	if err = usrAccount.SetEmail(req.Email); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if err = s.userRepo.SaveUserAccount(ctx, usrAccount); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) UpdatePhone(ctx context.Context, req *user.UpdatePhoneRequest) (resp *user.UpdatePhoneResponse, err error) {
	resp = new(user.UpdatePhoneResponse)

	usrAccount, err := s.userRepo.FindUserAccountByUserIDNonNil(ctx, req.UserId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	if err = usrAccount.SetPhoneNumber(req.PhoneNumber); err != nil {
		hlog.Infof("err %v", err)
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if err = s.userRepo.SaveUserAccount(ctx, usrAccount); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) UpdatePasswd(ctx context.Context, req *user.UpdatePasswdRequest) (resp *user.UpdatePasswdResponse, err error) {
	resp = new(user.UpdatePasswdResponse)

	usrAccount, err := s.userRepo.FindUserAccountByUserIDNonNil(ctx, req.UserId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	if usrAccount.Passwd != s.encryptManager.EncryptPassword(req.OldPasswd) {
		resp.BaseResp = utils.BuildBaseResp(errno.UserPwdError)
		return resp, nil
	}
	if req.OldPasswd == req.NewPasswd_ {
		resp.BaseResp = utils.BuildBaseResp(errno.UserPwdSameError)
		return resp, nil
	}

	usrAccount.SetPasswd(s.encryptManager.EncryptPassword(req.NewPasswd_))

	if err = s.userRepo.SaveUserAccount(ctx, usrAccount); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) UpdateBirthDay(ctx context.Context, req *user.UpdateBirthDayRequest) (resp *user.UpdateBirthDayResponse, err error) {
	resp = new(user.UpdateBirthDayResponse)

	usrProfile, err := s.userRepo.FindUserProfileByUserIDNonNil(ctx, req.UserId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	if err = usrProfile.SetBirthDay(req.Year, req.Month, req.Day); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	if err = s.userRepo.SaveUserProfile(ctx, usrProfile); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (resp *user.UploadAvatarResponse, err error) {
	resp = new(user.UploadAvatarResponse)

	_, err = s.userRepo.FindUserInfoByUserIDNonNil(ctx, req.UserId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	blobResp, err := s.blobManager.UploadBlob(ctx, &blob.GeneratePutPreSignedUrlRequest{
		UserId:   req.UserId,
		BlobType: consts.AvatarBlobType,
		Timeout:  int32(10 * time.Second.Seconds()),
	})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	resp.AvatarUrl = blobResp.Url
	resp.AvatarId = blobResp.Id
	resp.ObjectName = blobResp.ObjectName
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) UpdateAvatarInfo(ctx context.Context, req *user.UpdateAvatarInfoRequest) (resp *user.UpdateAvatarInfoResponse, err error) {
	resp = new(user.UpdateAvatarInfoResponse)

	usrInfo, err := s.userRepo.FindUserInfoByUserIDNonNil(ctx, req.UserId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, nil
	}

	_, err = s.blobManager.NotifyBlobUpload(ctx, &blob.NotifyBlobUploadRequest{
		BlobId:     req.AvatarId,
		UserId:     req.UserId,
		ObjectName: req.ObjectName,
		BlobType:   req.BlobType,
	})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	usrInfo.SetAvatarID(req.AvatarId)
	if err = s.userRepo.SaveUserInfo(ctx, usrInfo); err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserSrvError)
		return resp, nil
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) GetAvatar(ctx context.Context, req *user.GetAvatarRequest) (resp *user.GetAvatarResponse, err error) {
	resp = new(user.GetAvatarResponse)

	value := cache.AvatarItem{}
	avatarId := ""
	if err = s.cacheManager.Get(ctx, fmt.Sprintf(consts.BlobRedisKey, req.UserId, consts.AvatarBlobType), &value); err == nil && value.IsDirty() {
		avatarId = value.BlobID
	} else {
		usrInfo, err := s.userRepo.FindUserInfoByUserIDNonNil(ctx, req.UserId)
		if err != nil {
			resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
			return resp, nil
		}
		avatarId = usrInfo.AvatarID
		go func() {
			if err := s.cacheManager.Set(ctx, fmt.Sprintf(consts.BlobRedisKey, req.UserId, consts.AvatarBlobType),
				&cache.AvatarItem{BlobID: avatarId}); err != nil {
				klog.Infof("write cache error: %v", err)
			}
		}()
	}

	blobResp, err := s.blobManager.GetBlob(ctx, &blob.GenerateGetPreSignedUrlRequest{
		BlobId:  avatarId,
		Timeout: int32(60 * time.Second.Seconds()),
	})
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.BlobSrvError)
		return resp, nil
	}

	resp.AvatarUrl = blobResp.Url
	resp.BaseResp = utils.BuildBaseResp(nil)
	return resp, nil
}

func (s *UserApplicationService) CheckUserExist(ctx context.Context, req *user.CheckUserExistRequest) (resp *user.CheckUserExistResponse, err error) {
	resp = new(user.CheckUserExistResponse)

	userAccount, err := s.userRepo.FindUserAccountByUserID(ctx, req.UserId)
	if err != nil {
		resp.BaseResp = utils.BuildBaseResp(errno.UserNotExistError)
		return resp, err
	}

	resp.BaseResp = utils.BuildBaseResp(nil)
	if userAccount == nil {
		return resp, nil
	}
	resp.IsExist = true
	return resp, nil
}
