package runtime

import (
	"testing"
	"github.com/job-center/core/util/logger"
	"fmt"
)

func TestJob_NewJob(t *testing.T) {
	var job Job
	job.logger.LogRegister(logger.LevelDebug)
	job.NewJob(RestTrigger, tete)
}

func tete(args ... interface{}) {
	a := 5
	c := a + b
	fmt.Println(c)
}
