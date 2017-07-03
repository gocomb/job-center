package schedule

import "net/http"
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
type InitVariable struct {
}
type Rest struct {
	Method  string
	Path    string
	Name    string
	Handler http.Handler
}
type RpcApi struct {
}
