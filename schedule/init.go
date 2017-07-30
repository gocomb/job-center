package schedule

import (
	"net/http"
	"github.com/job-schedule/trigger/rest"
	"reflect"
	"github.com/job-schedule/message"
)
func initJobMesaage()*message.JobMessage{
	return &message.JobMessage{
		InfoMessage:  make(map[string]string),
		FatalMessage: make(map[string]error),
		ErrMessage:   make(map[string]string),
		WarnMessage:  make(map[string]string),
		DebugMassage: make(map[string]string),
	}
}

func initMessageChan() *RegisterParameter {
	return &RegisterParameter{
		FatalMessageChan: make(map[string]chan error),
		ErrMessageChan:   make(map[string]chan string),
		WarnMessageChan:  make(map[string]chan string),
		InfoMessageChan:  make(map[string]chan string),
		DebugMassageChan: make(map[string]chan string),
	}
}

func (job *Job) NewJob(trigger TriggerTypes, runTimeArg RegisterParameter, MyJob func(RegisterParameter, ...interface{}), jobArgs ...interface{}) {
	var err chan error
	err = make(chan error)
	job.Message =*(initJobMesaage())
	job.JobRunTimeArgs=*(initMessageChan())
	job.logger.SetInfo("Init trigger")
	go (job.initGetTriggerType(trigger))(&err)
	switch trigger {
	case RestTrigger:
		job.jobOneCycle.JobFunc = job.JobRunTime.JobFunc
	}
	go MyJob(runTimeArg, jobArgs ...)
	sind := reflect.Indirect(reflect.ValueOf(err))

	switch sind.Kind() {
	case reflect.Struct.String(), reflect.Slice:
		if sind.Len() == 0 {

		}
	}
	//监听监听端口发来的错误信息，若发生错误，线程结束。
	//todo: 暂用匿名函数临时表示，以后改成独立函数方法
	func(err *chan error) {

		trggerErrMessge := <-*err
		job.logger.SetInfo("trggerErrMessge", trggerErrMessge)
	}(&err)
}

func (c *JobRunTime) initGetTriggerType(trigger TriggerTypes) (func(restErr *chan error)) {
	//多种触发方式的初始化
	switch trigger {
	case EtcdTrigger:
	case RestTrigger:
		c.logger.SetInfo("Into the Rest Trigger")
		return c.restTriggerInit
	case RpcTrigger:

	}
	return c.restTriggerInit
}

//初始化路由，监听8080端口，若发生错误将错误信息传给chan，线程结束
//todo：rest的端口监听逻辑上应该不在job内部
func (c *JobRunTime) restTriggerInit(restErr *chan error) () {
	go c.logger.SetFatal(func(restErr *chan error) error {
		err := http.ListenAndServe(":8080", func() http.Handler {
			s, _ := rest.CollectRouters()
			return s
		}())
		if err != nil {
			c.logger.SetDebug("err is ", err)
			*restErr <- err
		} else {
			c.logger.SetInfo("Init the rest trigger is done")
		}
		return err
	}(restErr))
	return
}

func GetTriggerTypeDone() {

}

//job守护程序，控制job完成周期，阻塞与中断线程，daemon收集本job所有的message，将其传递给message
//todo：对job进行逻辑判断，以控制job其他协程的同步，当所状态结束后，daemon协程结束，job的生命周期结束
func (this *JobRunTime) jobDaemon() {
	//轮询message
	go this.getFatalChan()
}

