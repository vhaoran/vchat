package g

import (
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

var xUUIDget = new(uuidGet)

type uuidGet struct {
	sync.Mutex
}

func (r *uuidGet) Get() string {
	r.Lock()
	defer r.Unlock()

	count := 1
REDO:
	if s := getUUID(); len(s) > 0 {
		return s
	}

	if count++; count > 3 {
		return ""
	}
	time.Sleep(time.Microsecond * 50)
	goto REDO
}

//uuid "github.com/satori/go.uuid"
func GetUUID() string {
	return xUUIDget.Get()
}

func getUUID() string {
	//u, err := uuid.NewV1()
	u := uuid.NewV1()
	//if err != nil {
	//	return ""
	//}
	//
	return u.String()
}
