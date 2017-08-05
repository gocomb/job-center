package server

import (
	"context"

	"sync"

	"github.com/job-center/server/trigger/rest"

	"github.com/job-center/server/util"
)

func (i InitJob) Init() {
	if i.TriggerTypes == RestTrigger {
		go rest.RunRest()
	}
}

//run a job
//It is a the entrance of a job
func Run(jobName RegisterJob) {
	var threadCount sync.WaitGroup
	threadCount.Add(2)
	job := jobFactory(jobName)
	go job.jobRun(job.jobDone, &threadCount)
	go job.jobDaemon(job.jobDone, &threadCount)
	threadCount.Wait()
	//<-job.syncVar.ctxMessage

}

//register a job with specified parameters
func jobFactory(jobName RegisterJob) (r *jobServer) {
	r = new(jobServer)
	r.jobLog = make(map[string]string)
	r.RegisterJob = jobName
	r.syncVar = syncVar{jobDone: make(chan string),
		ctxMessage: make(chan string)}
	return r
}

//jobDaemon manage the job of all life cycle
func (j *jobServer) jobDaemon(done chan string, threadCount *sync.WaitGroup) {
	util.Logger.SetInfo("jobDaemon is running ")
	messageCtx, MessageCancel := context.WithCancel(context.Background())
	go util.Tick("second", "messageDaemon is running", j.messageDaemon, MessageCancel)
	<-done
	util.Logger.SetInfo("jobDaemon:job is done")
	go j.registerApiServer()
	MessageCancel()
	j.waitForQuit(messageCtx)
	util.Logger.SetInfo("jobDaemon done")
	threadCount.Done()
	util.Logger.SetInfo("wait for main done")
}

func (j *jobServer) registerApiServer() {

}

func (j *jobServer) waitForQuit(ctx context.Context) {
	<-ctx.Done()
}

//run a job
func (j *jobServer) jobRun(done chan string, threadCount *sync.WaitGroup) {
	j.NewJob(j.JobContext)
	done <- "jo"
	threadCount.Done()
	util.Logger.SetDebug("job server:job is done")
}
