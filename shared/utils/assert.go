package utils

import (
	"log"
	"os"
)

func Assert(cond bool, message string) {
	if !cond {
		log.Fatal(message)
	}
}

func AssertErr(err error) {
	if err != nil {
		// log.Fatal(fmt.Sprintf("%v\n%s", message, err))
		log.Fatal(err)
		os.Exit(1)
	}
}
