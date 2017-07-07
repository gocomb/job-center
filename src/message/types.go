package message

import "util/logger"

func init() {

}

type Interface interface {
	Info(...interface{}) func(parameter ...interface{})
}

type methodMessage struct {
}


//消息文本，输出消息信息，返回文本数组
func (c *methodMessage) Info(m ...interface{})interface{} {
	s := new(logger.NewLog)
	s.LogRegister(logger.LevelInfo)
	s.SetInfo(m...)
	return m
}
