package core

import (
	"github.com/job-schedule/core/message"
	"context"
)

type RegisterJob struct {
	NewJob     func(JobContext)
	JobMessage message.JobMessage
	ctx        context.Context
}
type JobContext struct {
	jobLog map[string]string
	ctx    context.Context
}

type jobServer struct {
	RegisterJob
	MessageChan
}

type JobMessage struct {
}
