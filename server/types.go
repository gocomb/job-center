package server

import (
	"github.com/job-center/pkg/message"
	"context"
)

type TriggerTypes int

const (
	_           TriggerTypes = iota
	EtcdTrigger
	RestTrigger
	RpcTrigger
)

type InitJob struct {
	TriggerTypes TriggerTypes
}
//Register a Job. It is the The first thing to run a job
// It can be simply as server.RegisterJob{NewJob: jobExample}
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
	syncVar
}

// jobServer is the base struct of job-centre
type jobServer struct {
	RegisterJob
	JobContext
	jobMessage
}
type jobMessage struct {
	messagePool message.Pool
	JobMessageChan
}

// JobMessageChan is a bridge link to pool and jobServer
type JobMessageChan struct {
	FatalMessageChan map[string]chan string
	ErrMessageChan   map[string]chan string
	WarnMessageChan  map[string]chan string
	InfoMessageChan  map[string]chan string
	DebugMassageChan map[string]chan string
}

//manage job life cycle
type syncVar struct {
	jobDone    chan string
	ctxMessage chan string
}
