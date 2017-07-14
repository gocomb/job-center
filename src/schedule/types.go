package schedule

import (
	"net/http"
	"sync"
	"util/logger"
	"message"
)

type TriggerTypes int

const (
	_           TriggerTypes = iota
	EtcdTrigger
	RestTrigger
	RpcTrigger
)

type Job struct {
	JobSchedule
	Trigger
	JobRunTime
}
type JobRunTime struct {
	jobOneCycle
	Message message.JobMessage
}
type JobSchedule struct {
	Trigger
}
type Trigger struct {
	RestApi       map[string]Rest
	RpcApi
	EtcdWatch
	ThreadTrigger ThreadTriggerTypes
}
type jobOneCycle struct {
	JobFunc func(interface{}) error
	FuncArgs interface{}
	initJob
}
type EtcdWatch struct {
}
type Rest struct {
	Method  string
	Path    string
	Name    string
	Handler http.Handler
}
type RpcApi struct {
}

type ThreadTriggerTypes struct {
	ThreadCount      map[string]sync.WaitGroup
	statusSwitchOn   chan bool
	statusSwitchOff  chan bool
	statusSwitchLast *bool
}

type initJob struct {
	ThreadCount      sync.WaitGroup
	statusSwitchOn   chan bool
	statusSwitchOff  chan bool
	errTrigger       chan error
	statusSwitchLast *bool
	LoggerLevel      logger.LogLevel
	logger           logger.NewLog
	TriggerType      TriggerTypes
	JobRunTimeArgs   RegisterParameter
}
type RegisterParameter struct {
	FatalMessageChan map[string]chan error
	ErrMessageChan   map[string]chan string
	WarnMessageChan  map[string]chan string
	InfoMessageChan  map[string]chan string
	DebugMassageChan map[string]chan string
}
