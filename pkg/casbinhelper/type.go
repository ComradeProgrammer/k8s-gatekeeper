package casbinhelper

import (
	"fmt"
	"reflect"
)

func IsNil(args ...interface{}) (interface{}, error) {
	fmt.Println(args[0])
	if len(args) != 1 {
		return nil, fmt.Errorf("IsNil requires 1 parameters, currently %d", len(args))
	}
	v := reflect.ValueOf(args[0])
	if v.Kind() == reflect.Pointer {
		return v.IsNil(), nil
	}
	return false, nil
}
