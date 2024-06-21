package entity

import (
	"github.com/google/uuid"
	"github.com/xince-fun/InstaGo/server/services/user/domain/vo"
	"github.com/xince-fun/InstaGo/server/shared/utils"
)

type UserProfile struct {
	UserID   uuid.UUID
	Bio      string
	BirthDay utils.LocalDate
	Gender   int
}

func NewUserProfile(userID uuid.UUID) *UserProfile {
	return &UserProfile{
		UserID: userID,
	}
}

func (u *UserProfile) SetUserID(userID uuid.UUID) {
	u.UserID = userID
}

func (u *UserProfile) SetBio(bio string) {
	u.Bio = bio
}

func (u *UserProfile) SetBirthDay(year, month, day int32) error {
	d, err := vo.NewDate(int(year), int(month), int(day))
	if err != nil {
		return err
	}
	u.BirthDay = d
	return nil
}

func (u *UserProfile) SetGender(gender int) {
	u.Gender = gender
}
