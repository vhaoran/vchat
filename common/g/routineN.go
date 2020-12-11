package g

import (
	"log"
	"sync"
	"time"
)

type (
	WaitGroupN struct {
		c  chan bool
		wg sync.WaitGroup

		sync.RWMutex
		start time.Time
	}
)

func NewWaitGroupN(count int) *WaitGroupN {
	bean := &WaitGroupN{
		c:     make(chan bool, count),
		start: time.Now(),
	}

	return bean
}

func (r *WaitGroupN) Call(f func() error) {
	r.c <- true
	r.wg.Add(1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
			//pop
			<-r.c
			r.wg.Done()
		}()
		f()
	}()
}

func (r *WaitGroupN) Wait() time.Duration {
	r.wg.Wait()

	r.Lock()
	defer r.Unlock()
	close(r.c)

	return time.Since(r.start)
}
