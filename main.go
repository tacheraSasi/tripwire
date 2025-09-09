package main

import (
	"fmt"
	"os"

	errorshandler "github.com/tacheraSasi/tripwire/errorsHandler"
	"github.com/tacheraSasi/tripwire/types"
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

	// OS checks
	fmt.Println("Is Linux:", utils.IsLinux())
	fmt.Println("Is Mac:", utils.IsMac())
	fmt.Println("Is Windows:", utils.IsWindows())

	// Random utilities
	fmt.Println("Random int [1,100]:", utils.GenerateRandInt(1, 100))
	fmt.Println("Random float [0,1]:", utils.GenerateRandFloat(0, 1))
	fmt.Println("Random string (8 chars):", utils.RandomString(8))

	// File/dir existence
	fmt.Println("LICENSE file exists:", utils.FileExists("LICENSE"))
	fmt.Println("Current dir exists:", utils.DirExists("."))

	// String and slice utilities
	fmt.Println("JoinStrings:", utils.JoinStrings(",", []string{"a", "b", "c"}))
	fmt.Println("SplitAndTrim:", utils.SplitAndTrim("a, b, c", ","))
	fmt.Println("PadLeft:", utils.PadLeft("42", 5, '0'))
	fmt.Println("PadRight:", utils.PadRight("go", 4, '-'))
	fmt.Println("UniqueStrings:", utils.UniqueStrings([]string{"a", "b", "a"}))
	fmt.Println("RemoveEmptyStrings:", utils.RemoveEmptyStrings([]string{"a", "", "b"}))
	fmt.Println("RepeatString:", utils.RepeatString("go", 3))
	fmt.Println("TruncateString:", utils.TruncateString("hello world", 5))
	fmt.Println("ChunkSlice:", utils.ChunkSlice([]int{1, 2, 3, 4, 5}, 2))

	// Math utilities
	fmt.Println("Clamp 10 to [0,5]:", utils.Clamp(10, 0, 5))
	fmt.Println("Min(3,7):", utils.Min(3, 7))
	fmt.Println("Max(3.2,7.1):", utils.Max(3.2, 7.1))
	fmt.Println("Abs(-5):", utils.Abs(-5))

	// ContainsAny
	fmt.Println("ContainsAny:", utils.ContainsAny("hello world", []string{"foo", "world"}))


	// Types package usage
	// --- Normal union ---
	s := types.New("hello")
	n := types.New(42)
	f := types.New(3.14)
	b := types.New(true)

	fmt.Println("String:", s.String())
	fmt.Println("Int:", n.String())
	fmt.Println("Float:", f.String())
	fmt.Println("Bool:", b.String())

	// --- Optional / nullable union ---
	var maybeName types.Union[string] = types.None[string]()
	if !maybeName.IsSome() {
		fmt.Println("No name set")
	}

	maybeName = types.New("Tach")
	if name, ok := maybeName.Get(); ok {
		fmt.Println("Name:", name)
	}
}
