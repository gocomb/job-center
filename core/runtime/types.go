package runtime

import (
	"net/http"
	"sync"
	"github.com/job-schedule/core/util/logger"
	"github.com/job-schedule/core/message"
)

type TriggerTypes int

const (
	_           TriggerTypes = iota
	EtcdTrigger
	RestTrigger
	RpcTrigger
)

type Job struct {
	JobRegister
	Trigger
	JobRunTime
}

type JobRunTime struct {
	LoggerLevel      logger.LogLevel
	logger           logger.NewLog
	Message message.JobMessage
	ThreadCount      sync.WaitGroup
	statusSwitchOn   chan bool
	statusSwitchOff  chan bool
	errTrigger       chan error
	statusSwitchLast *bool

}
type JobRegister struct {
	TriggerType      TriggerTypes
	JobRunTimeArgs   RegisterParameter
	JobFunc func(interface{}) error
	FuncArgs interface{}
}
type Trigger struct {
	RestApi       map[string]Rest
	RpcApi
	EtcdWatch
	ThreadTrigger ThreadTriggerTypes
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


type RegisterParameter struct {
	FatalMessageChan map[string]chan error
	ErrMessageChan   map[string]chan string
	WarnMessageChan  map[string]chan string
	InfoMessageChan  map[string]chan string
	DebugMassageChan map[string]chan string
}
