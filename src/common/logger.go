package common

import (
	"log"
	"net/http"
	"time"
)

var Debug0 = true
var Debug1 = false

func init() {
	if Debug1 == true {
		Debug0 = true
	}
	if Debug0 == false {
		Debug1 = false
	}
}
func DebugPrint(parameter ...interface{}) {
	if Debug0 == true {
		log.Println(parameter ...)
	}
}
func HttpLog(inner http.Handler, name string) http.Handler {
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
