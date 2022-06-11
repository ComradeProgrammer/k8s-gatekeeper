package casbinhelper

import (
	"fmt"
	"strings"
)

func HasPrefix(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("HasPrefix requires 2 parameters, currently %d", len(args))
	}
	str, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("HasPrefix requires 1st parameter to be string")
	}
	prefix, ok := args[1].(string)
	if !ok {
		return nil, fmt.Errorf("HasPrefix requires 2nd parameter to be string")
	}
	return strings.HasPrefix(str, prefix), nil
}

func Split(args ...interface{})(interface{},error){
	if len(args)<3{
		return nil, fmt.Errorf("split requires 3 parameters, currently %d", len(args))
	}
	sep,ok:=args[len(args)-2].(string)
	if !ok{
		return nil, fmt.Errorf("split requires penultimate 2nd parameters to be string")
	}
	posfloat,ok:=args[len(args)-1].(float64)
	if !ok{
		return nil, fmt.Errorf("split requires penultimate 1st parameters to be number")

	}
	pos:=int(posfloat)

	var res=make([]interface{},0)
	for i:=0;i<len(args)-2;i++{
		str,ok:=args[i].(string)
		if !ok{
			return nil, fmt.Errorf("split requires 1st parameters to be string")
		}
		splits:=strings.Split(str,sep)
		if len(splits)<=pos{
			return nil, fmt.Errorf("index overflow on string %s",str)
		}
		res=append(res, splits[pos])
	}
	return res,nil
}
