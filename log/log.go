package log

import (
	"log"
)

func ServerLog(err error) {
	log.SetPrefix("[Server] ")
	log.Println(err)
}
