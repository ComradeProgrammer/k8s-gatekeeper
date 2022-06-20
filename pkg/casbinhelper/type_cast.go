package casbinhelper

import (
	"fmt"
	"reflect"
)

func ToString(args ...interface{}) (interface{}, error) {
	fmt.Println("here",args)
	if len(args) != 1 {
		return nil, fmt.Errorf("ToString requires 1 parameters, currently %d", len(args))
	}
	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.String {
		return nil, fmt.Errorf("ToString: args[0] cannot be converted to string")

	}
	return v.String(), nil
}
