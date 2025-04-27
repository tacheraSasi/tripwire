package errorshandler

import "github.com/tacheraSasi/tripwire.git/utils"

func Check(err error, msg string) {
	if err != nil {
		utils.Print(msg)
		panic(err)
	}
}