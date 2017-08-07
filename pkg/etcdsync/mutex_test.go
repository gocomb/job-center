package etcdsync

import (
	"testing"

	"github.com/job-center/server/util"
)

func TestMutex_Lock(t *testing.T) {

	m := New("/etcdsync", 123, []string{"http://10.110.18.107:2379"})
	if m == nil {
		util.Logger.SetInfo("etcdsync.NewMutex failed")
	}
	err := m.Lock()
	if err != nil {
		util.Logger.SetInfo("etcdsync.Lock failed")
	} else {
		util.Logger.SetInfo("etcdsync.Lock OK")
	}

	util.Logger.SetInfo("Get the lock. Do something here.")

	err = m.Unlock()
	if err != nil {
		util.Logger.SetInfo("etcdsync.Unlock failed")
	} else {
		util.Logger.SetInfo("etcdsync.Unlock OK")
	}
}
