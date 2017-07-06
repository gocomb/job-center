package logger

import (
	"log"
	"net/http"
	"time"
)

func (a *NewLog) newLog() {
	switch a.loggerInit.LogLevel {
	case level1:
		a.Set1 = printOut
		a.Set2 = noneOut
		a.Set3 = noneOut
	case level2:
		a.Set1 = printOut
		a.Set2 = printOut
		a.Set3 = noneOut
	case level3:
		a.Set1 = printOut
		a.Set2 = printOut
		a.Set3 = printOut
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
