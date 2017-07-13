package message

import (
	"testing"
	"fmt"
)

func TestNewMessage_Info(t *testing.T) {
	c:=new(JobMessage)
	bb:=c.Info("sdf","asdasdas")
	fmt.Println(bb)
}