package main

import (
	"runtime"
	"time"
	"math/rand"

	"fmt"
	"github.com/job-center/core"
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
	core.Run(myServer)

}
