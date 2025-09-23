package utils

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"strings"
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

// IsLinux returns true if the current OS is Linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac returns true if the current OS is macOS.
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsWindows returns true if the current OS is Windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// GetEnvOrDefault returns the value of the environment variable or a default if not set.
func GetEnvOrDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

// FileExists returns true if the file exists and is not a directory.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// DirExists returns true if the directory exists.
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// JoinStrings joins a slice of strings with a separator.
func JoinStrings(sep string, elems []string) string {
	return strings.Join(elems, sep)
}

// SplitAndTrim splits a string by a separator and trims whitespace from each part.
func SplitAndTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	return parts
}

// PadLeft pads a string on the left to the specified length with padChar.
func PadLeft(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(padChar), length-len(s)) + s
}

// PadRight pads a string on the right to the specified length with padChar.
func PadRight(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(string(padChar), length-len(s))
}

// GenerateRandInt returns a random integer between min and max (inclusive).
func GenerateRandInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

// GenerateRandFloat returns a random float64 between min and max.
func GenerateRandFloat(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return min + rand.Float64()*(max-min)
}

// Min returns the minimum of two values.
func Min[T ~int | ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values.
func Max[T ~int | ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value.
func Abs[T ~int | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// ContainsAny returns true if any of the substrings are in s.
func ContainsAny(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// RemoveEmptyStrings removes empty strings from a slice.
func RemoveEmptyStrings(input []string) []string {
	result := make([]string, 0, len(input))
	for _, v := range input {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

// RepeatString repeats a string n times.
func RepeatString(s string, n int) string {
	return strings.Repeat(s, n)
}

// TruncateString truncates a string to maxLen, adding "..." if truncated.
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// ChunkSlice splits a slice into chunks of the given size.
func ChunkSlice[T any](s []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	var chunks [][]T
	for size < len(s) {
		chunks = append(chunks, s[:size])
		s = s[size:]
	}
	chunks = append(chunks, s)
	return chunks
}

// GetOS returns the current operating system as a string: "windows", "linux", or "mac".
func GetOS() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "mac"
	case "linux":
		return "linux"
	case "windows":
		return "windows"
	default:
		return "unknown"
	}
}


// GetInput prompts the user for input and returns the trimmed string.
func GetInput(promptText string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promptText, " ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// AskForConfirmation prompts the user with a yes/no question and returns true for 'y' and false for 'n' or default.
func AskForConfirmation(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/N]: ", prompt)
	resp, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(resp)) == "y"
}