package etcdsync

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	NewClient([]string{"http://10.110.18.107:2379"})
	c := GetClient()
	c.CreateDir("/demo")
	fmt.Println("demo is ", c.IsDirExist("/demo"))
	c.Set("demoKey", "demo val")
	val, err := c.Get("demoKey")
	fmt.Println("demo is ", val, err)
	c.Delete("demoKey")
}
