package persistence

import (
	"context"
	"github.com/google/uuid"
	"github.com/xince-fun/InstaGo/server/services/user/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/user/infra/persistence/dal"
	"github.com/xince-fun/InstaGo/server/shared/errno"
)

func NewUserRepo(userDal *dal.UserDal) *UserRepo {
	return &UserRepo{
		userDal: userDal,
	}
}

type UserRepo struct {
	userDal *dal.UserDal
}

func (r *UserRepo) NextIdentity() (uuid.UUID, error) {
	return uuid.NewV7()
}

func (r *UserRepo) SaveUser(ctx context.Context, user *entity.User) (err error) {
	tx := r.userDal.BeginTransaction(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	if err = r.userDal.UpsertAccount(ctx, user.UserAccount, tx); err != nil {
		return err
	}
	if err = r.userDal.UpsertInfo(ctx, user.UserInfo, tx); err != nil {
		return err
	}
	if err = r.userDal.UpsertProfile(ctx, user.UserProfile, tx); err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) SaveUserAccount(ctx context.Context, account *entity.UserAccount) error {
	return r.userDal.UpsertAccount(ctx, account, nil)
}

func (r *UserRepo) SaveUserProfile(ctx context.Context, profile *entity.UserProfile) error {
	return r.userDal.UpsertProfile(ctx, profile, nil)
}

func (r *UserRepo) SaveUserInfo(ctx context.Context, info *entity.UserInfo) error {
	return r.userDal.UpsertInfo(ctx, info, nil)
}

func (r *UserRepo) FindUserAccountByUserID(ctx context.Context, userID string) (*entity.UserAccount, error) {
	return r.userDal.SelectUserAccountByUserID(ctx, userID, nil)
}

func (r *UserRepo) FindUserProfileByUserID(ctx context.Context, userID string) (*entity.UserProfile, error) {
	return r.userDal.SelectUserProfileByUserID(ctx, userID, nil)
}

func (r *UserRepo) FindUserInfoByUserID(ctx context.Context, userID string) (*entity.UserInfo, error) {
	return r.userDal.SelectUserInfoByUserID(ctx, userID, nil)
}

func (r *UserRepo) FindUserAccountByUserIDNonNil(ctx context.Context, s string) (*entity.UserAccount, error) {
	rlt, err := r.FindUserAccountByUserID(ctx, s)
	if err != nil {
		return nil, err
	}
	if rlt == nil {
		return nil, errno.RecordNotFound
	}
	return rlt, nil
}

func (r *UserRepo) FindUserProfileByUserIDNonNil(ctx context.Context, s string) (*entity.UserProfile, error) {
	rlt, err := r.FindUserProfileByUserID(ctx, s)
	if err != nil {
		return nil, err
	}
	if rlt == nil {
		return nil, errno.RecordNotFound
	}
	return rlt, nil
}

func (r *UserRepo) FindUserInfoByUserIDNonNil(ctx context.Context, s string) (*entity.UserInfo, error) {
	rlt, err := r.FindUserInfoByUserID(ctx, s)
	if err != nil {
		return nil, err
	}
	if rlt == nil {
		return nil, errno.RecordNotFound
	}
	return rlt, nil
}

func (r *UserRepo) FindUserAccountByPhoneNumber(ctx context.Context, phone string) (*entity.UserAccount, error) {
	return r.userDal.SelectUserAccountByPhone(ctx, phone, nil)
}

func (r *UserRepo) FindUserAccountByPhoneNumberNonNil(ctx context.Context, s string) (*entity.UserAccount, error) {
	rlt, err := r.FindUserAccountByPhoneNumber(ctx, s)
	if err != nil {
		return nil, err
	}
	if rlt == nil {
		return nil, errno.RecordNotFound
	}
	return rlt, nil
}

func (r *UserRepo) FindUserAccountByEmail(ctx context.Context, email string) (*entity.UserAccount, error) {
	return r.userDal.SelectUserAccountByEmail(ctx, email, nil)
}

func (r *UserRepo) FindUserAccountByEmailNonNil(ctx context.Context, email string) (*entity.UserAccount, error) {
	rlt, err := r.FindUserAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if rlt == nil {
		return nil, errno.RecordNotFound
	}
	return rlt, nil
}

func (r *UserRepo) RemoveUser(ctx context.Context, userID string) (err error) {
	tx := r.userDal.BeginTransaction(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	if err = r.userDal.SoftDeleteUserAccount(ctx, userID, tx); err != nil {
		return
	}
	if err = r.userDal.SoftDeleteUserProfile(ctx, userID, tx); err != nil {
		return
	}
	if err = r.userDal.SoftDeleteUserInfo(ctx, userID, tx); err != nil {
		return
	}
	return nil
}
