package schedule

import (
	"net/http"
	"sync"
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
type Job struct {
	JobFunc func() error
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
type InitJob struct {
	InitVariable
}