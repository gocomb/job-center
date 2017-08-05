package server

import (
	"bytes"
	"context"

	"github.com/job-center/pkg/monitor"
	"github.com/job-center/server/util"
)

//声明jobChanCtx，jobChanCtx向job pool传递消息。
//job pool 收到消息后jobChan通道关闭。
//Job工作通过setLog将job信息存储在job的map中，每收到一个job日志，jobChanCtx向job pool传送一条消息，同时清掉当前map
// messageDaemon 通过range map 方式，将val 和key push到job pool中
func (j *jobServer) messageDaemon(cancel context.CancelFunc) {
	util.Logger.SetDebug("into messageDaemon func")
	j.newMessageChan()
	j.messagePool.NewPool()
	j.collectMessage(j.jobLog)
	//cancel()
	//j.messageChanFactory(j.collectFatalMessage())

}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetFatal(message string) *monitor.Method {
	c.preSetMessage(message, "fatal")
	return &monitor.Method{Message: message, DoContext: c.ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetError(message string) *monitor.Method {
	c.preSetMessage(message, "error")
	return &monitor.Method{Message: message, DoContext: c.ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetWarn(message string) *monitor.Method {
	c.preSetMessage(message, "warn")
	return &monitor.Method{Message: message, DoContext: c.ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetDebug(message string) *monitor.Method {
	c.preSetMessage(message, "debug")
	return &monitor.Method{Message: message, DoContext: c.ctxMessage}
}

//SetFatal set a fatal message log, it return a monitor.Method struct, by which the method of message can be call
func (c *JobContext) SetInfo(message string) *monitor.Method {
	c.preSetMessage(message, "info")
	return &monitor.Method{Message: message, DoContext: c.ctxMessage}
}
func (c *JobContext) preSetMessage(message string, spec string) {
	c.jobLog[spec+util.GenTag(message)] = message
	util.Logger.SetDebug(c.jobLog)

}

//create the fatal chan  as the same number of all fatal message
//Chan is link message pool and in order to implement message queue
func (c *jobMessage) watchFatal() {

}

//collectMessage is the common method of collect message
//Chan is link message pool and in order to implement message queue
func (j *jobMessage) collectMessage(job map[string]string) {
	for k, v := range job {
		j.preSendMessage(k, v, "fatal")
		j.preSendMessage(k, v, "warn")
		j.preSendMessage(k, v, "debug")
		j.preSendMessage(k, v, "info")
		j.preSendMessage(k, v, "error")
	}
}

//prepare all fatal message from job to message pool
func (j *jobMessage) preSendMessage(key string, val string, match string) {
	if bytes.Contains([]byte(key), []byte(match)) {
		//j.FatalMessageChan[key] <- val
	}
}

// MessageChanFactory is declare a JobMessageChan variable pointer
func MessageChanFactory() JobMessageChan {
	return JobMessageChan{
		FatalMessageChan: make(map[string]chan string),
		ErrMessageChan:   make(map[string]chan string),
		WarnMessageChan:  make(map[string]chan string),
		InfoMessageChan:  make(map[string]chan string),
		DebugMassageChan: make(map[string]chan string),
	}
}

// newMessageChan is get a new JobMessageChan
func (j *jobMessage) newMessageChan() {
	j.JobMessageChan = MessageChanFactory()
}
