package g

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

//byte[] string
//singe value to string
//other to json
func GetBufferForMq(in interface{}) ([]byte, error) {
	var s []byte
	var err error
	v := reflect.Indirect(reflect.ValueOf(in))
	if !v.IsValid() {
		return nil, errors.New("in is null")
	}

	switch in.(type) {
	case string:
		return []byte(in.(string)), nil
	case []byte:
		return (in.([]byte)), nil
	}

	switch v.Type().Kind() {
	case reflect.Struct, reflect.Array:
		s, err = json.Marshal(in)
		if err != nil {
			return nil, err
		}
		return s, nil
	}

	str := fmt.Sprint(v)

	return []byte(str), nil
}
