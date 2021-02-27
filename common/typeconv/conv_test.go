package typeconv

import (
	"fmt"
	"testing"
)

func Test_aaa(t *testing.T) {
	//
	{
		s := "t"
		fmt.Println("as int:", NewStrResult(s).AsInt(2))
		fmt.Println("as int64:", NewStrResult(s).AsInt64(3))
		fmt.Println("as bool:", NewStrResult(s).AsBool(true))
		fmt.Println("as str:", NewStrResult(s).AsStr())

	}

}
