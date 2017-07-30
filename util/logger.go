package util

import mylogger"util/logger"
var Logger mylogger.NewLog
func init()  {
	Logger.LogRegister(mylogger.LevelDebug)
}
