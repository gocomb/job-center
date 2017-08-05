package runtime

import (
	"time"
)


func (a *jobOneCycle) runOneCycle() error {
	a.ThreadCount.Add(1)
	err := a.JobFunc(a.FuncArgs)
	return err
}
