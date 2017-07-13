package schedule

import (
	"time"
)


func (a *jobOneCycle) runOneCycle() error {
	a.ThreadCount.Add(1)
	err := a.JobFunc()
	return err
}
