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

func (u *UserInfo) SetAvatarID(avatarID string) {
	u.AvatarID = avatarID
}

func (u *UserInfo) GetID() string {
	return u.UserID.String()
}

func (u *UserInfo) IsDirty() bool {
	return u.FullName != "" || u.Account != "" || u.AvatarID != ""
}
