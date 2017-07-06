package schedule

import (
	"log"
	"net/http"
	"trigger/rest"
)

func (c *InitJob) initGetTriggerType() func() {
	switch c.InitVariable {
	case EtcdTrigger:
	case RestTrigger:
		return RestTriggerInit
	case RpcTrigger:

	}
	return GetTriggerTypeDone
}

func RestTriggerInit() {
	log.Fatal(http.ListenAndServe(":8080", func() http.Handler {
		s, _ := rest.CollectRouters()
		return s
	}()))
}

func GetTriggerTypeDone() {

}
