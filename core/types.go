package core

import (
	"github.com/job-schedule/core/message"
	"context"
)

//Register a Job. It is the The first thing to run a job
// It can be simply as core.RegisterJob{NewJob: jobExample}
type RegisterJob struct {
	NewJob     func(JobContext)
	JobMessage message.JobMessage
	ctx        context.Context
}

//JobContext is the struct that make job-centre get the job running status.
//Now the status is only by output log, which is define by JobMessage
//In order to handle message, the JobMessage would be classify to jobLog
//As the parameter of job, JobContext is a bridge link to jobServer
type JobContext struct {
	jobLog map[string]string
	ctx    context.Context
}

// jobServer is the base struct of job-centre
type jobServer struct {
	RegisterJob
	messagePool message.Pool
}

type JobMessage struct {
}
