package main

import (
	"errors"

	errorshandler "github.com/tacheraSasi/tripwire.git/errorsHandler"
)


func main() {
	err := errors.New("this is an error")
	errorshandler.Check(err,"an error occurred")
}