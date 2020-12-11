package ipLoction

import (
	"fmt"
	"log"
	"testing"
)

func Test_ip_get(t *testing.T) {
	bean, err := NewCurrentIP("47.104.24.37")
	log.Println(err)
	fmt.Println("------", "", "-----------")
	log.Println(bean)
}

