package cache

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	temp := Register("myRepo")
	temp.Put("asd", "asddd")
	isOk := temp.MatchQueryKey("as")
	fmt.Println(temp.Get("asd"), isOk)
}
