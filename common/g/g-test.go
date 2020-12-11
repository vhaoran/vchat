package g

import (
	"fmt"
	"time"
)

func RandStr() string {
	return fmt.Sprint("name_", time.Now().UnixNano())
}

func RandInt64() int64 {
	return time.Now().UnixNano()
}

func RandInt32() int32 {
	return int32(time.Now().UnixNano() % int64(1<<31))
}
