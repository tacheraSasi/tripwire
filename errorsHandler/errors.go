package errorshandler

import (
	"os"

	"github.com/tacheraSasi/tripwire/utils"
)

func Check(err error, msg string) {
	if err != nil {
		utils.Print(msg, err)
		os.Exit(1)
	}
}

func CheckNoExit(err error, msg string) {
	if err != nil {
		utils.Print(msg, err)
	}
}

func Panic(err error, msg string) {
	if err != nil {
		utils.Print(msg, err)
		panic(err)
	}
}