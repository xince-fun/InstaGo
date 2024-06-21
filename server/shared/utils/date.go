package utils

import (
	"database/sql/driver"
	"fmt"
	"github.com/xince-fun/InstaGo/server/shared/consts"
	"time"
)

type LocalTime time.Time

func (t LocalTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return []byte(time.Time(t).Format(consts.TimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("LocalDate: unable to scan value of type %T", v)
	}
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", value.Format(consts.TimeFormat))
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(consts.TimeFormat)
}

type LocalDate time.Time

func (t LocalDate) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return []byte(time.Time(t).Format(consts.DateFormat)), nil
}

func (t *LocalDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("LocalDate: unable to scan value of type %T", v)
	}
	tTime, _ := time.Parse("2006-01-02", value.Format(consts.DateFormat))
	*t = LocalDate(tTime)
	return nil
}

func (t LocalDate) String() string {
	return time.Time(t).Format(consts.DateFormat)
}
