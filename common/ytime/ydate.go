package ytime

import (
	"database/sql/driver"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"log"
	"time"
)

const (
	//TimeZone ...
	TimeZone = "Asia/Shanghai"
	//CustomDateFmt ...
	CustomDateFmt = "2006-01-02 15:04:05"

	//DateLayout ...
	DateLayout = "2006-01-02"
)

func init() {
	SetTimeZone()
	log.Println("-----ytime init-called---------")
}

//SetTimeZone ...
func SetTimeZone() {
	lc, err := time.LoadLocation(TimeZone)
	if err == nil {
		time.Local = lc
	}
}

//Date ...
type Date struct {
	time.Time
}

//根据当天日期返回一个0时的时间值
func OfToday() Date {
	return Date{time.Now()}.ToDate()
}

//返回一个含时、分、秒的时间值
func OfNow() Date {
	return Date{Time: time.Now()}
}

//OfDatetime ...
func OfStr(in string) (Date, error) {
	out, err := time.ParseInLocation(CustomDateFmt, in, time.Local)
	return Date{out}, err
}

func OfTime(t time.Time) Date {
	return Date{t}
}

func OfInt(year, month, day int, l ...int) Date {
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
	bean, err := OfStr(s)
	if err != nil {
		return OfNow()
	}
	return bean
}

//String ...
func (p Date) ToDate() Date {
	return OfInt(p.Year(), int(p.Month()), p.Day())
}

//String ...
func (p Date) ToStr() string {
	return p.Format(CustomDateFmt)
}

func (p Date) ToStrShort() string {
	return p.Format("2006-01-02_15_04_05")
}

func (p Date) ToStrDate() string {
	return p.Format("2006-01-02")
}

func (t Date) MarshalJSON() ([]byte, error) {
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

//UnmarshalJSON ...
func (p *Date) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"`+CustomDateFmt+`"`, string(data), time.Local)
	if err != nil {
		*p = Date{Time: time.Now()}
	}
	*p = Date{Time: local}
	return nil
}

// Value insert timestamp into mysql need this function.
func (p Date) Value() (driver.Value, error) {
	var zeroTime time.Time
	if p.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return p.Time, nil
}

// Scan value of time.Time
func (p *Date) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*p = Date{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (p Date) TimeShanghai() time.Time {
	return p.Time.In(time.Local)
}

//------------------------------
// UnmarshalBSON unmarshal bson
func (p *Date) UnmarshalBSON(data []byte) (err error) {
	fmt.Println("#### Un-MarshalBSON ", p, " #####")

	var d bson.D
	err = bson.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	if v, ok := d.Map()["time"]; ok {
		p.Time = time.Time{}
		return p.Time.UnmarshalText([]byte(v.(string)))
	}
	return fmt.Errorf("key 't' missing")
}

// MarshalBSON marshal bson
func (p Date) MarshalBSON() ([]byte, error) {
	fmt.Println("#### MarshalBSON ", p, " #####")

	txt, err := p.Time.MarshalText()
	if err != nil {
		return nil, err
	}
	b, err := bson.Marshal(map[string]string{"time": string(txt)})
	return b, err
}

// MarshalBSONValue marshal bson value
func (p *Date) MarshalBSONValue() (bsontype.Type, []byte, error) {
	fmt.Println("#### MarshalBSONValue ", p, " #####")
	b, err := bson.Marshal(p)
	return bson.TypeEmbeddedDocument, b, err
}
