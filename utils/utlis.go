package utils

import (
	"fmt"
	"io"
)

// Print prints the provided arguments to standard output.
var Print = fmt.Println

// Ternary returns trueVal if condition is true, otherwise falseVal.
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// StringInSlice returns true if the target string is found in the slice.
func StringInSlice(target string, list []string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

// ReverseSlice reverses a slice of any type in place.
func ReverseSlice[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// SafeClose closes an io.Closer and ignores nil, returning any error.
func SafeClose(c io.Closer) error {
	if c != nil {
		return c.Close()
	}
	return nil
}
