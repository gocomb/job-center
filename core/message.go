package core

import (
	"github.com/job-schedule/core/monitor"
	"github.com/job-schedule/core/util"

)

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetFatal(message string) *monitor.Method {
	c.jobLog["fatal"+util.GenTag(message)] = message
	util.Logger.SetDebug(c.jobLog)
	return &monitor.Method{Message: message, DoContext: ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetError(message string) *monitor.Method {
	c.jobLog["error"+util.GenTag(message)] = message
	util.Logger.SetDebug(c.jobLog)
	return &monitor.Method{Message: message, DoContext: ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetWarn(message string) *monitor.Method {
	c.jobLog["warn"+util.GenTag(message)] = message
	util.Logger.SetDebug(c.jobLog)
	return &monitor.Method{Message: message, DoContext: ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetDebug(message string) *monitor.Method {
	c.jobLog["debug"+util.GenTag(message)] = message
	util.Logger.SetDebug(c.jobLog)
	return &monitor.Method{Message: message, DoContext: ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetInfo(message string) *monitor.Method {
	c.jobLog["info"+util.GenTag(message)] = message
	util.Logger.SetDebug(c.jobLog)
	return &monitor.Method{Message: message, DoContext: ctxMessage}
}

//collect all fatal message from job
func (c *JobContext) collectFatal() {
	collectMessage(c.jobLog)
}

//create the fatal chan  as the same number of all fatal message
//Chan is link message pool and in order to implement message queue
func (c *JobContext) watchFatal() {

}

//collectMessage is the common method of collect message
//Chan is link message pool and in order to implement message queue
func collectMessage(job map[string]string) {
	for k, v := range job {
		util.Logger.SetInfo(k, v)
	}
}
