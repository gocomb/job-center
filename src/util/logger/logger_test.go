package logger

import "testing"

func TestGenLog_OutL1(t *testing.T)  {
	var s NewLog
	s.LogRegister(level1)
	s.Set3("asd","sdfsdf")
}