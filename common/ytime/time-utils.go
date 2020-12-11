package ytime

import (
	"time"
)

func Today() time.Time {
	t := time.Now()
	return DayTime(t)
}

func DayTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func HourTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local)
}

func MinuteTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
		t.Minute(),
		0,
		0, time.Local)
}

func SecondTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
		t.Minute(),
		t.Second(),
		0, time.Local)
}

//下月的1号0:0:0
func NextMonth(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), 1
	//12月时调整
	if m == 12 {
		y, m = y+1, 1
	} else {
		m++
	}

	return OfInt(y, int(m), d, 0, 0, 0).Time
}

//上一月的1号
func PriorMonth(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), 1
	//12月时调整
	if m == 1 {
		y, m = y-1, 12
	} else {
		m--
	}

	return OfInt(y, int(m), d, 0, 0, 0).Time
}

//本月的1号
func CurrentMonth(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), 1
	//12月时调整
	return OfInt(y, int(m), d, 0, 0, 0).Time
}
