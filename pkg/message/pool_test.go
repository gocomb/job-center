package message

import (
	"fmt"
	"testing"
)

func TestPool_FetchAllFatalMessage(t *testing.T) {
	var poolRepoTest Pool
	poolRepoTest.NewPool()
	poolRepoTest.PushMessage("fatal1234456", "hello my fatal1")
	poolRepoTest.PushMessage("fatal12344567", "hello my fatal2")
	fmt.Println(poolRepoTest.FetchAllFatalMessage())
}
