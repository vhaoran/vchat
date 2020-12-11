package g

import (
	"time"
)

// to hour/minute/second to 0
func ToDate(src time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	y, m, d := src.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}
