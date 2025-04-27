package main

import (
	"os"

	errorshandler "github.com/tacheraSasi/tripwire.git/errorsHandler"
)


func main() {
	filename := "nonexistent.txt"
	file, err := os.Open(filename)
	errorshandler.Panic(err,"an error occurred while opening the file"+filename)
	defer file.Close()
}