package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/job-center/server"
	"github.com/job-center/server/util"
)

func jobExample(ctx server.JobContext) {
	util.Logger.SetDebug("hello,jobExample is start")
	util.Logger.SetDebug("now, set a fatal message 1")
	ctx.SetFatal("hello fatal message").Do("print")
	util.Logger.SetDebug("now, set a info message")
	ctx.SetInfo("hello info message").Do("print")
	util.Logger.SetDebug("now, set a warn message")
	ctx.SetWarn("hello warn message").Do("print")
	util.Logger.SetDebug("now, set a debug message")
	ctx.SetError("hello error message").Do("print")
	util.Logger.SetDebug("now, set a error message")
	ctx.SetDebug("hello debug message").Do("print")
	util.Logger.SetDebug("now, set a fatal message 2")
	ctx.SetFatal("hello fatal message x").Do("print")
	util.Logger.SetDebug("now wait for 5s")
	time.Sleep(5000 * time.Millisecond)
	util.Logger.SetDebug("jobExample is done")


}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	a:=server.InitJob{server.RestTrigger}
	a.Init()
	myServer := server.RegisterJob{NewJob: jobExample}
	server.Run(myServer)

}
