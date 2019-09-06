package errors

import (
	"log"
)

func CheckError(err error, text string) {
	if err != nil {
		log.Print(text)
		log.Print(err)
	}
}
