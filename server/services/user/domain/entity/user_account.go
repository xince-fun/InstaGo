package entity

import (
	"errors"
	"github.com/google/uuid"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"github.com/xince-fun/InstaGo/server/shared/utils"
)

type UserAccount struct {
	UserID      uuid.UUID
	Email       string
	PhoneNumber string
	Passwd      string
}

func NewUserAccount(userID uuid.UUID) *UserAccount {
	return &UserAccount{
		UserID: userID,
	}
}

func (u *UserAccount) SetUserID(userID uuid.UUID) {
	u.UserID = userID
}

func (u *UserAccount) SetPasswd(passwd string) {
	u.Passwd = passwd
}

func (u *UserAccount) SetPhoneNumber(phoneNumber string) error {
	if valid := utils.IsValidRegexp(consts.PhoneNumberRegexp, phoneNumber); !valid {
		return errors.New("phoneNumber format error")
	}
	if phoneNumber != "" {
		u.PhoneNumber = phoneNumber
	}
	return nil
}

func (u *UserAccount) SetEmail(email string) error {
	if valid := utils.IsValidRegexp(consts.EmailRegexp, email); !valid {
		return errors.New("email format error")
	}
	if email != "" {
		u.Email = email
	}
	return nil
}
