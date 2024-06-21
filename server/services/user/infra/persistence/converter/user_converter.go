package converter

import (
	"github.com/xince-fun/InstaGo/server/services/user/domain/entity"
	"github.com/xince-fun/InstaGo/server/services/user/infra/persistence/po"
)

func UserAccountToPo(userAccount *entity.UserAccount) *po.UserAccount {
	return &po.UserAccount{
		UserID:      userAccount.UserID,
		Email:       userAccount.Email,
		PhoneNumber: userAccount.PhoneNumber,
		Passwd:      userAccount.Passwd,
	}
}

func UserAccountToEntity(userAccount *po.UserAccount) *entity.UserAccount {
	return &entity.UserAccount{
		UserID:      userAccount.UserID,
		Email:       userAccount.Email,
		PhoneNumber: userAccount.PhoneNumber,
		Passwd:      userAccount.Passwd,
	}
}

func UserProfileToPo(userProfile *entity.UserProfile) *po.UserProfile {
	return &po.UserProfile{
		UserID:   userProfile.UserID,
		Bio:      userProfile.Bio,
		BirthDay: userProfile.BirthDay,
		Gender:   userProfile.Gender,
	}
}

func UserProfileToEntity(userProfile *po.UserProfile) *entity.UserProfile {
	return &entity.UserProfile{
		UserID:   userProfile.UserID,
		Bio:      userProfile.Bio,
		BirthDay: userProfile.BirthDay,
		Gender:   userProfile.Gender,
	}
}

func UserInfoToPo(userInfo *entity.UserInfo) *po.UserInfo {
	return &po.UserInfo{
		UserID:   userInfo.UserID,
		Account:  userInfo.Account,
		FullName: userInfo.FullName,
		AvatarID: userInfo.AvatarID,
	}
}

func UserInfoToEntity(userInfo *po.UserInfo) *entity.UserInfo {
	return &entity.UserInfo{
		UserID:   userInfo.UserID,
		Account:  userInfo.Account,
		FullName: userInfo.FullName,
		AvatarID: userInfo.AvatarID,
	}
}
