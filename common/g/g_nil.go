package g

import (
	"fmt"
	"reflect"
)

//传入的bean应该能返回str,Sprintf("%v",bean)
func NilStr(bean interface{}, defaultStr string) string {
	v := reflect.Indirect(reflect.ValueOf(bean))

	//
	if !v.IsValid() {
		return defaultStr
	}

	//

	//------------对于str,只考虑----------------
	switch v.Kind() {
	case reflect.String:
		if len(v.String()) == 0 {
			return defaultStr
		}
	}

	s := fmt.Sprintf("%v", bean)
	return s
}
