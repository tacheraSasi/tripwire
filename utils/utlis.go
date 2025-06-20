package utils

import (
	"fmt"
	"io"
	"math/rand"
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

// MapSlice applies a function to each element of a slice and returns a new slice with the results.
func MapSlice[T any, U any](s []T, fn func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

// FilterSlice returns a new slice containing only the elements that satisfy the predicate.
func FilterSlice[T any](s []T, fn func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// UniqueStrings returns a new slice with duplicate strings removed.
func UniqueStrings(input []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(input))
	for _, v := range input {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// RandomString generates a random string of the given length using letters and digits.
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// IntInSlice returns true if the target int is found in the slice.
func IntInSlice(target int, list []int) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

// Clamp restricts x to the range [min, max].
func Clamp[T ~int | ~float64](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
