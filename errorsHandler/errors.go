package errorshandler

import (
	"os"

	"github.com/tacheraSasi/tripwire/utils"
)

// Check exits the program if err is not nil, printing the provided message and error.
func Check(err error, msg string) {
	if err != nil {
		utils.Print("[FATAL]", msg+":", err)
		os.Exit(1)
	}
}

// CheckNoExit prints the error and message if err is not nil, but does not exit.
func CheckNoExit(err error, msg string) {
	if err != nil {
		utils.Print("[ERROR]", msg+":", err)
	}
}

// Panic panics if err is not nil, printing the provided message and error.
func Panic(err error, msg string) {
	if err != nil {
		utils.Print("[PANIC]", msg+":", err)
		panic(err)
	}
}
