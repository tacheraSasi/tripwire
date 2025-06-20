# Tripwire

[![Go Report Card](https://goreportcard.com/badge/github.com/tacheraSasi/tripwire)](https://goreportcard.com/report/github.com/tacheraSasi/tripwire)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Tripwire is a simple Go utility library for error handling and utility functions, designed for extensibility and open source collaboration.

## Features
- Idiomatic error handling helpers
- Generic utility functions (e.g., ternary)
- Easy to use and extend

## Installation

```
go get github.com/tacheraSasi/tripwire
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/tacheraSasi/tripwire/errorsHandler"
    "github.com/tacheraSasi/tripwire/utils"
)

func main() {
    file, err := os.Open("LICENSE")
    errorsHandler.Check(err, "Failed to open file")
    defer file.Close()
    fmt.Println(utils.Ternary(true, "yes", "no"))
}
```

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License
MIT. See [LICENSE](LICENSE).
