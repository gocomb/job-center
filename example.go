package main

import (
	"runtime"
	"time"
	"math/rand"

	"fmt"
	"github.com/job-schedule/core"
)

func jobExample(ctx core.JobContext) {
	fmt.Println("hello")
	ctx.SetFatal("hello fatal message").Do("print")
	ctx.SetFatal("hello fatal message x").Do("print")


}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	myServer := core.RegisterJob{NewJob: jobExample}
	myServer.JobMessage.FatalMessage = make(map[string]error)
	core.Run(myServer)

	/*go func() {

	}()
	log.Fatal(http.ListenAndServe(":8080", func() http.Handler {
		s, _ := rest.CollectRouters()
		return s
	}()))*/
}
