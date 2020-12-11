package g

import (
	"fmt"
	"log"
	"testing"
)

func Test_routineN(t *testing.T) {
	obj := NewWaitGroupN(20)

	h := 100000
	for i := 0; i < h; i++ {
		j := i
		obj.Call(func() error {
			log.Println(j, "-->")
			if j%100 == 1 {
				panic(fmt.Sprint("err", j))
			}
			return nil
		})
	}
	obj.Wait()
}

func Test_g_in(t *testing.T) {
	b := In(1, 1, 2, 3)
	fmt.Println("b = ", b)
}
