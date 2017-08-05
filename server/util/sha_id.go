package util

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func GenTag(info string) string {
	data := []byte(info)
	tag := sha1.Sum(data)
	return fmt.Sprintf("%x", tag)

}
func GenTagRandom(info string) string {
	data := []byte(info + time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"))
	tag := sha1.Sum(data)
	return fmt.Sprintf("%x", tag)

}
