package ytime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

//Date ...
type DateM time.Time

//根据当天日期返回一个0时的时间值
func OfTodayM() DateM {
	return DateM(time.Now()).ToDate()
}

//返回一个含时、分、秒的时间值
func OfNowM() DateM {
	return DateM(time.Now())
}

//OfDatetime ...
func OfStrM(in string) (DateM, error) {
	out, err := time.ParseInLocation(CustomDateFmt, in, time.Local)
	return DateM(out), err
}

func OfTimeM(t time.Time) DateM {
	return DateM(t)
}

func OfIntM(year, month, day int, l ...int) DateM {
	hour, minute, second := 0, 0, 0
	if len(l) > 0 {
		hour = l[0]
	}
	if len(l) > 1 {
		minute = l[1]
	}
	if len(l) > 2 {
		second = l[2]
	}
	s := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	bean, err := OfStrM(s)
	if err != nil {
		return OfNowM()
	}
	return bean
}

//String ...
func (p DateM) ToDate() DateM {
	return OfIntM(time.Time(p).Year(),
		int(time.Time(p).Month()),
		time.Time(p).Day())
}

//String ...
func (p DateM) ToStr() string {
	return time.Time(p).Format(CustomDateFmt)
}

func (p DateM) ToStrShort() string {
	return time.Time(p).Format("2006-01-02_15_04_05")
}

func (p DateM) ToStrDate() string {
	return time.Time(p).Format("2006-01-02")
}

func (p DateM) MarshalJSON() ([]byte, error) {
	//tune := (time.Time(p)).Format(`"2006-01-02 15:04:05"`)
	//return []byte(tune), nil
	if y := time.Time(p).Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	t := time.Time(p)

	b := make([]byte, 0, len(CustomDateFmt)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, CustomDateFmt)
	b = append(b, '"')
	return b, nil
}

//UnmarshalJSON ...
func (p *DateM) UnmarshalJSON(data []byte) error {
	t, err := OfStrM(string(data))
	if err == nil {
		*p = t
	}
	return err
}

// Value insert timestamp into mysql need this function.
func (p DateM) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(p).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(p), nil
}

// Scan value of time.Time
func (p *DateM) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*p = DateM(value)
		return nil
	}
	//
	v1, ok1 := v.(DateM)
	if ok1 {
		*p = v1
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (p DateM) TimeShanghai() time.Time {
	return time.Time(p).In(time.Local)
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (p DateM) MarshalText() ([]byte, error) {
	str := p.ToStr()
	return []byte(str), nil
}
