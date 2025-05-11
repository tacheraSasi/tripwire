package utils

import "fmt"

var Print = fmt.Println

// Ternary operations in golang
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}
