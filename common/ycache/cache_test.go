package ycache

import (
	"fmt"
	"testing"
	"time"
)

func Test_cac(t *testing.T) {
	//
	c := NewCacheExpire(
	10, time.Second)
	//----------------------------------------------
	for i := 0; i < 3; i++ {
		v, _ := c.GetIfPresent(i)
		fmt.Println("get: ", v)
	}
	fmt.Println("---------get 2-------------")
	for i := 0; i < 3; i++ {
		v, _ := c.GetIfPresent(i)
		fmt.Println("get: ", v)
	}

	//----------------------------------------------
	time.Sleep(time.Millisecond * 900)
	fmt.Println(" after sleep.....")
	for i := 0; i < 10; i++ {
		_, _ = c.GetIfPresent(i)
		time.Sleep(100 * time.Millisecond)
	}
}
