# Tripwire

[![Go Report Card](https://goreportcard.com/badge/github.com/tacheraSasi/tripwire)](https://goreportcard.com/report/github.com/tacheraSasi/tripwire)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Tripwire is a simple Go utility library for error handling and utility functions, designed for extensibility and open source collaboration.

## Features
- Idiomatic error handling helpers
- Generic utility functions (e.g., ternary, slice helpers, random string)
- Easy to use and extend

## Installation

```
go get github.com/tacheraSasi/tripwire
```

## Usage

Import the packages:

```go
import (
    "github.com/tacheraSasi/tripwire/errorsHandler"
    "github.com/tacheraSasi/tripwire/utils"
)
```

### Error Handling

- **Check**: Exits the program if error is not nil.
  ```go
  errorsHandler.Check(err, "Failed to open file")
  ```
- **CheckNoExit**: Prints error if not nil, but does not exit.
  ```go
  errorsHandler.CheckNoExit(err, "Non-fatal error")
  ```
- **Panic**: Panics if error is not nil.
  ```go
  errorsHandler.Panic(err, "Critical error")
  ```

### Utilities

- **Print**: Shortcut for fmt.Println.
  ```go
  utils.Print("Hello, world!")
  ```
- **Ternary**: Generic ternary operator.
  ```go
  result := utils.Ternary(x > 0, "positive", "non-positive")
  ```
- **StringInSlice**: Check if a string is in a slice.
  ```go
  found := utils.StringInSlice("foo", []string{"bar", "foo"}) // true
  ```
- **IntInSlice**: Check if an int is in a slice.
  ```go
  found := utils.IntInSlice(3, []int{1, 2, 3}) // true
  ```
- **ReverseSlice**: Reverse a slice in place.
  ```go
  nums := []int{1, 2, 3}
  utils.ReverseSlice(nums) // nums is now [3, 2, 1]
  ```
- **MapSlice**: Map a function over a slice.
  ```go
  squares := utils.MapSlice([]int{1, 2, 3}, func(x int) int { return x * x }) // [1, 4, 9]
  ```
- **FilterSlice**: Filter a slice by a predicate.
  ```go
  evens := utils.FilterSlice([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }) // [2, 4]
  ```
- **UniqueStrings**: Remove duplicate strings from a slice.
  ```go
  unique := utils.UniqueStrings([]string{"a", "b", "a"}) // ["a", "b"]
  ```
- **RandomString**: Generate a random alphanumeric string.
  ```go
  s := utils.RandomString(8) // e.g., "aZ3kLm2Q"
  ```
- **SafeClose**: Safely close an io.Closer, handling nil.
  ```go
  utils.SafeClose(file)
  ```
- **Clamp**: Clamp a value between min and max.
  ```go
  v := utils.Clamp(10, 0, 5) // v == 5
  v2 := utils.Clamp(-1.5, 0.0, 1.0) // v2 == 0.0
  ```

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License
MIT. See [LICENSE](LICENSE).
