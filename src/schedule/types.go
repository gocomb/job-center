package types

import (
	"net/http"
	"sync"
)
var ThreadCount sync.WaitGroup
var statusSwitchOn chan bool
var statusSwitchOff chan bool
var statusSwitchLast *bool
type JobSchedule struct {
	Trigger
	Job
	InitVariable
}
type Trigger struct {
	RestApi map[string]Rest
	RpcApi
}
type Job struct {
	JobFunc func()error
}

type Rest struct {
	Method  string
	Path    string
	Name    string
	Handler http.Handler
}
type RpcApi struct {
}
