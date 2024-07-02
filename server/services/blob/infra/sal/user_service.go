package sal

import (
	"context"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/blob/pkg/initialize"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user"
	"github.com/xince-fun/InstaGo/server/shared/kitex_gen/user/userservice"
	"sync"
)

var UserManagerSet = wire.NewSet(
	NewUserManager,
)

type UserManager struct {
	client userservice.Client
	once   sync.Once
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (u *UserManager) InitClient() {
	u.once.Do(func() {
		u.client = initialize.InitUser()
	})
}

func (u *UserManager) UpdateAvatarInfo(ctx context.Context, req *user.UpdateAvatarInfoRequest) (resp *user.UpdateAvatarInfoResponse, err error) {
	u.InitClient()
	return u.client.UpdateAvatarInfo(ctx, req)
}
