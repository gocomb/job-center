package schedule

import (
	"github.com/job-schedule/core/util/logger"
	"runtime"
)

//传输fatal信息通道，daemon收到fatal信号后结束线程
func (j *JobRunTime) getFatalChan() {
	for k, v := range j.JobRunTimeArgs.FatalMessageChan {
		go j.catchFatalMessage(j.Message.FatalMessage[k], v)
	}
}
func (j *JobRunTime) catchFatalMessage(err error, errChan chan error) {
	select {
	case message := <-errChan:
		err = message
		if j.LoggerLevel == logger.LevelDebug {
			//todo:退出主协程
			runtime.Goexit()
		}
	}
}

//传输其他信息通道接口，daemon收到后发送message
func (j *JobRunTime) getMessageChan() {
	for k, v := range j.JobRunTimeArgs.InfoMessageChan {
		go func() {
			select {
			case message := <-v:
				j.Message.InfoMessage[k] = message

			}
		}()
	}
}
