package core

import (
	"context"
	"reflect"
	"fmt"
	"github.com/job-center/core/util"
)

var ctxMessage chan string = make(chan string)

//run a job
//It is a the entrance of a job
func Run(jobName RegisterJob) {
	var (
		done chan string = make(chan string)
		job  *jobServer  = new(jobServer)
		ctx              = JobContext{}
	)

	job = job.jobFactory(jobName)
	//初始化map

	ctx.jobLog = make(map[string]string)
	//初始化结束

	go job.NewJob(ctx)
	go job.jobDaemon(ctx, done)
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

//Generate a log object based on the log level
func (j *jobServer) logFactory() {

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
func (j *jobServer) collectFatalMessage() map[string]string {
	return j.JobMessage.FatalMessage
}
func (j *jobServer) initMessage() {

}
