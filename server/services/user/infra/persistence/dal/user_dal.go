package dal

import (
	"context"
	"errors"
	"github.com/xince-fun/InstaGo/server/services/user/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/user/infra/persistence/converter"
	"github.com/xince-fun/InstaGo/server/services/user/infra/persistence/po"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func NewUserDal(db *gorm.DB) *UserDal {
	return &UserDal{db: db}
}

type UserDal struct {
	db *gorm.DB
}

func (d *UserDal) UpsertAccount(ctx context.Context, userAccount *entity.UserAccount, tx *gorm.DB) (err error) {
	userAccountPo := converter.UserAccountToPo(userAccount)

	now := utils.LocalTime(time.Now())
	userAccountPo.UpdateTime = now
	userAccountPo.CreateTime = now
	if tx == nil {
		tx = d.db
	}

	err = tx.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"email", "phone_number", "passwd", "is_deleted",
			"update_time"}),
	}).Create(userAccountPo).Error

	return err
}

func (d *UserDal) UpsertProfile(ctx context.Context, userProfile *entity.UserProfile, tx *gorm.DB) (err error) {
	userProfilePo := converter.UserProfileToPo(userProfile)

	now := utils.LocalTime(time.Now())
	userProfilePo.UpdateTime = now
	userProfilePo.CreateTime = now

	if tx == nil {
		tx = d.db
	}

	err = tx.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"bio", "birth_day", "gender", "is_deleted",
			"update_time"}),
	}).Create(userProfilePo).Error

	return err
}

func (d *UserDal) UpsertInfo(ctx context.Context, userInfo *entity.UserInfo, tx *gorm.DB) (err error) {
	userInfoPo := converter.UserInfoToPo(userInfo)

	now := utils.LocalTime(time.Now())
	userInfoPo.UpdateTime = now
	userInfoPo.CreateTime = now

	if tx == nil {
		tx = d.db
	}

	err = tx.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"account", "full_name", "avatar_id",
			"is_deleted", "update_time"}),
	}).Create(userInfoPo).Error

	return err
}

func (d *UserDal) SelectUserAccountByUserID(ctx context.Context, userID string, tx *gorm.DB) (*entity.UserAccount, error) {
	userAccountPo := po.UserAccount{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"user_id": userID, "is_deleted": 0}).First(&userAccountPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	userAccount := converter.UserAccountToEntity(&userAccountPo)
	return userAccount, nil
}

func (d *UserDal) SelectUserProfileByUserID(ctx context.Context, userID string, tx *gorm.DB) (*entity.UserProfile, error) {
	userProfilePo := po.UserProfile{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"user_id": userID, "is_deleted": 0}).First(&userProfilePo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	userProfile := converter.UserProfileToEntity(&userProfilePo)
	return userProfile, nil
}

func (d *UserDal) SelectUserInfoByUserID(ctx context.Context, userID string, tx *gorm.DB) (*entity.UserInfo, error) {
	userInfoPo := po.UserInfo{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"user_id": userID, "is_deleted": 0}).First(&userInfoPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	userInfo := converter.UserInfoToEntity(&userInfoPo)
	return userInfo, nil
}

func (d *UserDal) SelectUserAccountByEmail(ctx context.Context, email string, tx *gorm.DB) (*entity.UserAccount, error) {
	userAccountPo := po.UserAccount{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"email": email, "is_deleted": 0}).First(&userAccountPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	userAccount := converter.UserAccountToEntity(&userAccountPo)
	return userAccount, nil
}

func (d *UserDal) SelectUserAccountByPhone(ctx context.Context, phone string, tx *gorm.DB) (*entity.UserAccount, error) {
	userAccountPo := po.UserAccount{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"phone_number": phone, "is_deleted": 0}).First(&userAccountPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	userAccount := converter.UserAccountToEntity(&userAccountPo)
	return userAccount, nil
}

func (d *UserDal) SoftDeleteUserAccount(ctx context.Context, userID string, tx *gorm.DB) error {
	userAccountPo := po.UserAccount{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"user_id": userID, "is_deleted": 0}).First(&userAccountPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return res.Error
	}

	userAccountPo.IsDeleted = 1
	userAccountPo.UpdateTime = utils.LocalTime(time.Now())

	return tx.WithContext(ctx).Model(&userAccountPo).Update("is_deleted", userAccountPo.IsDeleted).Error
}

func (d *UserDal) SoftDeleteUserProfile(ctx context.Context, userID string, tx *gorm.DB) error {
	userProfilePo := po.UserProfile{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"user_id": userID, "is_deleted": 0}).First(&userProfilePo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return res.Error
	}

	userProfilePo.IsDeleted = 1
	userProfilePo.UpdateTime = utils.LocalTime(time.Now())

	return tx.WithContext(ctx).Model(&userProfilePo).Update("is_deleted", userProfilePo.IsDeleted).Error
}

func (d *UserDal) SoftDeleteUserInfo(ctx context.Context, userID string, tx *gorm.DB) error {
	userInfoPo := po.UserInfo{}

	if tx == nil {
		tx = d.db
	}

	res := tx.WithContext(ctx).Where(map[string]interface{}{"user_id": userID, "is_deleted": 0}).First(&userInfoPo)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return res.Error
	}

	userInfoPo.IsDeleted = 1
	userInfoPo.UpdateTime = utils.LocalTime(time.Now())

	return tx.WithContext(ctx).Model(&userInfoPo).Update("is_deleted", userInfoPo.IsDeleted).Error
}

func (d *UserDal) BeginTransaction(ctx context.Context) *gorm.DB {
	return d.db.WithContext(ctx).Begin()
}
