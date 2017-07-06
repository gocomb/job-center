package schedule

import (
	"net/http"
	"sync"
	"util/logger"
)

type JobSchedule struct {
	Trigger
	Job
	Init InitJob
}
type Trigger struct {
	RestApi       map[string]Rest
	RpcApi
	EtcdWatch
	ThreadTrigger ThreadTriggerTypes
}
type jobOneCycle struct {
	JobFunc func() error
	InitVariable
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
type JobRunTime struct {
	jobOneCycle

}
type InitVariable struct {
	ThreadCount      sync.WaitGroup
	statusSwitchOn   chan bool
	statusSwitchOff  chan bool
	statusSwitchLast *bool
	LoggerLevel      logger.LogLevel
	logger           logger.NewLog
	TriggerType      TriggerTypes
}