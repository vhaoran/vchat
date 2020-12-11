package g

import (
	"fmt"
	"strings"
)

func IsDigitNoDot(s string) bool {
	if len(s) == 0 {
		return false
	}

	pat := "1234567890"
	for _, c := range s {
		sub := fmt.Sprintf("%c", c)
		if strings.Index(pat, sub) == -1 {
			return false
		}
	}
	//
	return true
}
