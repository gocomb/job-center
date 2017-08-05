package rest

import (
	"log"
	"net/http"
)

func RunRest() {
	log.Fatal(http.ListenAndServe(":8080", func() http.Handler {
		s, _ := CollectRouters()
		return s
	}()))
}
