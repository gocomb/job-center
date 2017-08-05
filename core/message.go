package core

import (
	"github.com/job-center/core/monitor"
	"github.com/job-center/core/util"
	"bytes"
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

//create the fatal chan  as the same number of all fatal message
//Chan is link message pool and in order to implement message queue
func (c *JobContext) watchFatal() {

}

//collectMessage is the common method of collect message
//Chan is link message pool and in order to implement message queue
func (j *JobMessageChan) collectMessage(job map[string]string) {
	for k, v := range job {
		j.sendMessageChan(k, v, "fatal")
		j.sendMessageChan(k, v, "warn")
		j.sendMessageChan(k, v, "debug")
		j.sendMessageChan(k, v, "info")
		j.sendMessageChan(k, v, "error")
	}
}

//prepare all fatal message from job to message pool
func (j *JobMessageChan) sendMessageChan(key string, val string, match string) {
	if bytes.Contains([]byte(key), []byte(match)) {
		j.FatalMessageChan[key] <- val
	}
}

// MessageChanFactory is declare a JobMessageChan variable pointer
func MessageChanFactory() *JobMessageChan {
	return &JobMessageChan{
		FatalMessageChan: make(map[string]chan string),
		ErrMessageChan:   make(map[string]chan string),
		WarnMessageChan:  make(map[string]chan string),
		InfoMessageChan:  make(map[string]chan string),
		DebugMassageChan: make(map[string]chan string),
	}
}
// newMessageChan is get a new JobMessageChan
func (j *JobMessageChan) newMessageChan() {
	j = MessageChanFactory()
}
