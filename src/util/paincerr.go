package common

import (
	"log"
	"time"
)

func LogErr(err error) {
	start := time.Now()
	if err != nil {
		log.Println("%s\t%s\t", time.Since(start), err)
	}

}
