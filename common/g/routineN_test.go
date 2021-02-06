package g

import (
	"fmt"
	"log"
	"testing"
	"time"
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

func Test_md5(t *testing.T) {
	//
	t0 := time.Now()
	h := 100000
	c := 0
	for i := 0; i < h; i++ {
		s := MD5(fmt.Sprint(i) + "_a")
		c += len(s)
		fmt.Println(s)
	}

	fmt.Println(c)
	fmt.Println(time.Now().Sub(t0).Milliseconds())
}

func Test_hash(t *testing.T) {
	//
	t0 := time.Now()
	h := 100
	c := 0
	for i := 0; i < h; i++ {
		s := Hash(fmt.Sprint(i) + "_a")
		c += len(s)
		fmt.Println(s)
		fmt.Println(len(s))
		fmt.Println(len(s))
	}

	fmt.Println(c)
	fmt.Println(time.Now().Sub(t0).Milliseconds())
}
