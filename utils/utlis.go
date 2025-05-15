package utils

import "fmt"

var Print = fmt.Println

func Ternary[T any](condition bool,trueVal, falseVale T)T{
	if condition{
		return trueVal;
	}
	return falseVale
}