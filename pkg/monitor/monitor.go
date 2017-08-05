package monitor

import (

	"github.com/job-center/server/util"
)
//handle the message method
type Method struct {
	Message string
	//todo: DoContext encapsulated into a normal context struct
	DoContext chan string
}

func NewMonitor() {

}

//the method function
func (c *Method) Do(method string) {
	util.Logger.SetDebug("hello monitor")
	c.doByContext(method, c.DoContext)

}

// manage the life cycle of method
func (c *Method) doByContext(method string, ctx chan string) {
	switch method {
	case "print":
		util.Logger.SetFatal(c.Message)
	}
	//by experiment, in golang, like xx.Func.Func is parallel process，needs to be processed synchronously
	//ctx <- "done"
}
