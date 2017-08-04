package message

import (
	. "github.com/job-center/core/cache"
)

type JobMessage struct {
	FatalMessage map[string]string
	ErrMessage   map[string]string
	WarnMessage  map[string]string
	InfoMessage  map[string]string
	DebugMassage map[string]string
	Context
}
type Context struct {
	FatalMessageChan map[string]chan string
	ErrMessageChan   map[string]chan string
	WarnMessageChan  map[string]chan string
	InfoMessageChan  map[string]chan string
	DebugMassageChan map[string]chan string
}
type Pool struct {
	//the cache of the message cache.
	Cache
}
