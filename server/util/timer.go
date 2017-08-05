package util

import (
	"context"
	"log"
	"strconv"
	"time"
)

func Tick(times string, string string, fun func(context.CancelFunc), cancel context.CancelFunc) {
	var ticker *time.Ticker
	switch times {
	case "minute":
		ticker = time.NewTicker(time.Minute)
	case "hour":
		ticker = time.NewTicker(time.Minute * 60)
	case "day":
		ticker = time.NewTicker(time.Minute * 60 * 24)
	case "second":
		ticker = time.NewTicker(time.Millisecond * 1000)
	case "monitor":
		ticker = time.NewTicker(time.Millisecond * 50)

	}

	i := 0
	for range ticker.C {
		i = i + 1
		fun(cancel)
		log.Printf("%s\t%s\t%s\t%s\t", string, "ticked at ", strconv.Itoa(i), time.Now())
	}

}
