package entity

import "github.com/google/uuid"

type UserInfo struct {
	UserID   uuid.UUID
	Account  string
	FullName string
	AvatarID string
}

func NewUserInfo(userID uuid.UUID) *UserInfo {
	return &UserInfo{
		UserID: userID,
	}
}

func (u *UserInfo) SetUserID(userID uuid.UUID) {
	u.UserID = userID
}

func (u *UserInfo) SetFullName(fullName string) {
	u.FullName = fullName
}

func (u *UserInfo) SetAccount(account string) {
	u.Account = account
}

func (u *UserInfo) SetAvatarUrl(avatarID string) {
	u.AvatarID = avatarID
}
