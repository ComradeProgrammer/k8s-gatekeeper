package casbinhelper

import (
	"fmt"
	"reflect"
)

func Len(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Index requires 2 parameters, currently %d", len(args))
	}
	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("1st parameter should be array, currently %s", v.Kind().String())
	}
	return v.Len(), nil

}
