package core

import (
	"context"
	"reflect"
	"fmt"
	"github.com/job-schedule/core/util"
	"github.com/job-schedule/core/monitor"
)

var ctxMessage chan string = make(chan string)

func (c *JobContext) SetFatal(message string) *monitor.Method {
	c.jobLog["fatal"+util.GenTag(message)] = message
	util.Logger.SetDebug(c.jobLog)
	return &monitor.Method{Message: message, DoContext: ctxMessage}
}

//run a job
func Run(jobName RegisterJob) {
	var done chan string
	done = make(chan string)
	ut := JobContext{}
	job := new(jobServer)
	job = job.jobFactory(jobName)
	//初始化map

	ut.jobLog = make(map[string]string)
	//初始化结束

	go job.NewJob(ut)
	go job.jobDaemon(ut, done)
	<-done
	<-ctxMessage

}

//register a job with specified parameters
func (r *jobServer) jobFactory(jobName RegisterJob) *jobServer {
	r.RegisterJob = jobName
	return r
}

//jobDaemon manage the job of all life cycle
func (j *jobServer) jobDaemon(ut JobContext, done chan string) {
	messageCtx, MessageCancel := context.WithCancel(context.Background())
	go j.messageDaemon(messageCtx)
	go j.registerApiServer()
	MessageCancel()
	j.waitForQuit(messageCtx)

	util.Logger.SetInfo("done")
	done <- "done"

}
func (j *jobServer) messageDaemon(ctx context.Context) {
	j.messageChanFactory(j.collectFatalMessage())

}

func (j *jobServer) registerApiServer() {

}

func (j *jobServer) waitForQuit(ctx context.Context) {
	<-ctx.Done()
}

func newJob() {

}

func getMapKey(MapChan interface{}) ([]reflect.Value) {
	return reflect.ValueOf(MapChan).MapKeys()

}
func (j *jobServer) messageChanFactory(message interface{}) {
	for k, v := range getMapKey(message) {
		fmt.Println(k, v)
	}
}
func (j *jobServer) collectFatalMessage() map[string]error {
	return j.JobMessage.FatalMessage
}
func (j *jobServer) initMessage() {

}
