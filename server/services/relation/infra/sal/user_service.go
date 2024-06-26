package sal

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user/userservice"
)

var UserManagerSet = wire.NewSet(
	NewUserManager,
)

type UserManager struct {
	client userservice.Client
}

func NewUserManager(client userservice.Client) *UserManager {
	return &UserManager{
		client: client,
	}
}

func (u *UserManager) CheckUserExist(ctx context.Context, req *user.CheckUserExistRequest) (*user.CheckUserExistResponse, error) {
	return u.client.CheckUserExist(ctx, req)
}

func (u *UserManager) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	return u.client.GetUserInfo(ctx, req)
}
