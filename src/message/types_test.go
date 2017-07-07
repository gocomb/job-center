package message

import (
	"testing"
	"fmt"
)

func TestNewMessage_Info(t *testing.T) {
	c:=new(NewMessage)
	bb:=c.Info("sdf","asdasdas")
	fmt.Println(bb)
}