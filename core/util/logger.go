package util

import mylogger"github.com/job-schedule/core/util/logger"
var Logger mylogger.NewLog
func init() {
	Logger.LogRegister(mylogger.LevelDebug)
}