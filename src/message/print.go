package message

import "util/logger"

func (c *JobMessage) Info(m ...interface{})interface{} {
	s := new(logger.NewLog)
	s.LogRegister(logger.LevelInfo)
	s.SetInfo(m...)
	return m
}

