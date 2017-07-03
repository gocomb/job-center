package main

import (
	"runtime"
	"time"
	"math/rand"
	"cmd/app"
	"log"
	"net/http"
	"trigger/rest"
	"util"
)

func main() {
	if util.Debug0 == true {
		util.DebugPrint("The cpu core is", runtime.NumCPU(), ",The app would use all of cores")
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	go app.Run()
	log.Fatal(http.ListenAndServe(":8080", func() http.Handler {
		s, _ :=  rest.CollectRouters()
		return s
	}()))
}
