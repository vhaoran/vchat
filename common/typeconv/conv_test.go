package typeconv

import (
	"fmt"
	"testing"
)

func Test_aaa(t *testing.T) {
	//
	{
		s := "t"
		fmt.Println("as int:", NewStrData(s).AsInt(2))
		fmt.Println("as int64:", NewStrData(s).AsInt64(3))
		fmt.Println("as bool:", NewStrData(s).AsBool(true))
		fmt.Println("as str:", NewStrData(s).AsStr())

	}

}
