package main

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/user/app"
	user "github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
)

var UserServiceImplSet = wire.NewSet(
	app.UserApplicationServiceSet,
	NewUserServiceImpl,
)

func NewUserServiceImpl(userAppService *app.UserApplicationService) *UserServiceImpl {
	return &UserServiceImpl{
		app: userAppService,
	}
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	app *app.UserApplicationService
}

// RegisterPhone implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterPhone(ctx context.Context, req *user.RegisterPhoneRequest) (resp *user.RegisterResponse, err error) {
	return s.app.RegisterPhone(ctx, req)
}

// RegisterEmail implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterEmail(ctx context.Context, req *user.RegisterEmailRequest) (resp *user.RegisterResponse, err error) {
	return s.app.RegisterEmail(ctx, req)
}

// LoginPhone implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginPhone(ctx context.Context, req *user.LoginPhoneRequest) (resp *user.LoginResponse, err error) {
	return s.app.LoginPhone(ctx, req)
}

// LoginEmail implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginEmail(ctx context.Context, req *user.LoginEmailRequest) (resp *user.LoginResponse, err error) {
	return s.app.LoginEmail(ctx, req)
}

// UpdateEmail implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateEmail(ctx context.Context, req *user.UpdateEmailRequest) (resp *user.UpdateEmailResponse, err error) {
	return s.app.UpdateEmail(ctx, req)
}

// UpdatePhone implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePhone(ctx context.Context, req *user.UpdatePhoneRequest) (resp *user.UpdatePhoneResponse, err error) {
	return s.app.UpdatePhone(ctx, req)
}

// UpdatePasswd implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePasswd(ctx context.Context, req *user.UpdatePasswdRequest) (resp *user.UpdatePasswdResponse, err error) {
	return s.app.UpdatePasswd(ctx, req)
}

// UpdateBirthDay implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateBirthDay(ctx context.Context, req *user.UpdateBirthDayRequest) (resp *user.UpdateBirthDayResponse, err error) {
	return s.app.UpdateBirthDay(ctx, req)
}

// UploadAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, req *user.UploadAvatarRequest) (resp *user.UploadAvatarResponse, err error) {
	return s.app.UploadAvatar(ctx, req)
}

// UpdateAvatarInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateAvatarInfo(ctx context.Context, req *user.UpdateAvatarInfoRequest) (resp *user.UpdateAvatarInfoResponse, err error) {
	return s.app.UpdateAvatarInfo(ctx, req)
}

// GetAvatar implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAvatar(ctx context.Context, req *user.GetAvatarRequest) (resp *user.GetAvatarResponse, err error) {
	return s.app.GetAvatar(ctx, req)
}
