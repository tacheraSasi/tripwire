package types

import "fmt"

// Extend this list with any types you want to allow
type UnionType interface {
	~string | ~int | ~float64 | ~bool
}

// --- Union wrapper ---
type Union[T UnionType] struct {
	val *T
}

// New creates a union with a concrete value
func New[T UnionType](v T) Union[T] {
	return Union[T]{val: &v}
}

// None creates a union with no value (nil)
func None[T UnionType]() Union[T] {
	return Union[T]{val: nil}
}

// IsSome checks if the union has a value
func (u Union[T]) IsSome() bool {
	return u.val != nil
}

// Get returns the stored value and a boolean indicating presence
func (u Union[T]) Get() (T, bool) {
	if u.val == nil {
		var zero T
		return zero, false
	}
	return *u.val, true
}

// MustGet returns the value or panics if nil
func (u Union[T]) MustGet() T {
	if u.val == nil {
		panic("Union has no value")
	}
	return *u.val
}

// --- Optional: helper to print type dynamically ---
func (u Union[T]) String() string {
	if u.val == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", *u.val)
}
