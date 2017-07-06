package schedule

import (
	"sync"
	"util/logger"
)


type TriggerTypes int

const (
	_           TriggerTypes = iota
	EtcdTrigger
	RestTrigger
	RpcTrigger
)
