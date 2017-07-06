package logger

import (
	"log"
	"net/http"
	"time"
)
//static Level DEBUG
//DEBUG Level指出细粒度信息事件对调试应用程序是非常有帮助的。
//static Level INFO
//INFO level表明 消息在粗粒度级别上突出强调应用程序的运行过程。
//static Level WARN
//WARN level表明会出现潜在错误的情形。
//static Level ERROR
//ERROR level指出虽然发生错误事件，但仍然不影响系统的继续运行。
//static Level FATAL
//FATAL level指出每个严重的错误事件将会导致应用程序的退出。
func (a *NewLog) newLog() {
	switch a.loggerInit.LogLevel {
	case LevelFatal:
		*a = NewLog{logInit{LevelDebug}, noneOut,
			noneOut, noneOut, noneOut,printOut}
	case LevelError:
		*a = NewLog{logInit{LevelInfo}, noneOut,
			noneOut, noneOut, printOut, printOut}
	case LevelWarn:
		*a = NewLog{logInit{LevelWarn}, noneOut,
			noneOut, printOut, printOut, printOut}
	case LevelInfo:
		*a = NewLog{logInit{LevelError}, noneOut,
			printOut, printOut, printOut, printOut}
	case LevelDebug:
		*a = NewLog{logInit{LevelFatal}, printOut,
			printOut, printOut, printOut, printOut}
	}
}

func (a *NewLog) LogRegister(c LogLevel) {

	a.loggerInit.LogLevel = c
	a.newLog()

}

func noneOut(none ...interface{}) {
}

func printOut(parameter ...interface{}) {
	log.Println(parameter...)
}

func (NewLog) HttpLog(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
