package main

import (
	"fmt"
	"os"

	errorshandler "github.com/tacheraSasi/tripwire/errorsHandler"
	"github.com/tacheraSasi/tripwire/utils"
)

// main is the entry point of the application.
func main() {
	filename := "LICENSE"
	file, err := os.Open(filename)
	output := utils.Ternary(err != nil, err.Error(), "File opened successfully")
	fmt.Println("Output:", output)
	errorshandler.Panic(err, "An error occurred while opening the file: "+filename)
	if file != nil {
		defer file.Close()
	}
}
