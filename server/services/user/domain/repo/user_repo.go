package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/xince-fun/InstaGo/server/services/user/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/user/infra/persistence"
	"github.com/xince-fun/InstaGo/server/services/user/infra/persistence/dal"
)

type UserRepository interface {
	NextIdentity() (uuid.UUID, error)

	SaveUser(context.Context, *entity.User) error
	SaveUserAccount(context.Context, *entity.UserAccount) error
	SaveUserProfile(context.Context, *entity.UserProfile) error
	SaveUserInfo(context.Context, *entity.UserInfo) error

	FindUserAccountByUserID(context.Context, string) (*entity.UserAccount, error)
	FindUserProfileByUserID(context.Context, string) (*entity.UserProfile, error)
	FindUserInfoByUserID(context.Context, string) (*entity.UserInfo, error)
	FindUserAccountByUserIDNonNil(context.Context, string) (*entity.UserAccount, error)
	FindUserProfileByUserIDNonNil(context.Context, string) (*entity.UserProfile, error)
	FindUserInfoByUserIDNonNil(context.Context, string) (*entity.UserInfo, error)
	FindUserAccountByPhoneNumber(context.Context, string) (*entity.UserAccount, error)
	FindUserAccountByPhoneNumberNonNil(context.Context, string) (*entity.UserAccount, error)
	FindUserAccountByEmail(context.Context, string) (*entity.UserAccount, error)
	FindUserAccountByEmailNonNil(context.Context, string) (*entity.UserAccount, error)
	RemoveUser(context.Context, string) error
}

var UserRepositorySet = wire.NewSet(
	dal.NewUserDal,
	persistence.NewUserRepo,
	wire.Bind(new(UserRepository), new(*persistence.UserRepo)),
)
