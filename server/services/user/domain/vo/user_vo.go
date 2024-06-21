package vo

import (
	"github.com/xince-fun/InstaGo/server/shared/errno"
	"github.com/xince-fun/InstaGo/server/shared/utils"
	"time"
)

func NewDate(year, month, day int) (utils.LocalDate, error) {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if t.Year() != year || t.Month() != time.Month(month) || t.Day() != day {
		return utils.LocalDate{}, errno.InvalidDate
	}
	return utils.LocalDate(t), nil
}
