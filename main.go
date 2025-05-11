package main

import (
	"fmt"
	"os"

	errorshandler "github.com/tacheraSasi/tripwire.git/errorsHandler"
	"github.com/tacheraSasi/tripwire.git/utils"
)

func main() {
	filename := "nonexistent.txt"
	file, err := os.Open(filename)
	output := utils.Ternary(err != nil, err, nil)
	fmt.Println("Output", output)
	errorshandler.Panic(err, "an error occurred while opening the file"+filename)
	defer file.Close()
}
