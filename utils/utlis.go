package utils

import "fmt"

// Print prints the provided arguments to standard output.
var Print = fmt.Println

// Ternary returns trueVal if condition is true, otherwise falseVal.
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
