package errors

import (
	"log"

	"github.com/matscus/mq-golang-jms20/jms20subset"
)

func CheckError(err error, str string) {
	if err != nil {
		log.Println(str, err)
	}
}
func CheckJMSError(err jms20subset.JMSException, str string) {
	if err != nil {
		log.Println(str, err)
	}
}
