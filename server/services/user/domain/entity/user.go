package entity

import (
	"github.com/google/uuid"
)

type User struct {
	UserID      uuid.UUID
	UserAccount *UserAccount
	UserProfile *UserProfile
	UserInfo    *UserInfo
}

func NewUser(userID uuid.UUID) *User {
	return &User{
		UserAccount: NewUserAccount(userID),
		UserProfile: NewUserProfile(userID),
		UserInfo:    NewUserInfo(userID),
	}
}

func (u *User) SetUserID(userID uuid.UUID) {
	u.UserID = userID
}
