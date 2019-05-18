package log

import (
	"log"
)

func ServerLog(err error) {
	log.SetPrefix("[Server] ")
	log.Println(err)
}

func ServerLogs(errs []error) {
	log.SetPrefix("[Server] ")
	for err := range errs {
		log.Println(err)
	}
}
