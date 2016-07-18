package ua

import (
	"fmt"
	"testing"
)

func TestDetect(t *testing.T) {
	cnt := 0
	errCnt := 0
	for _, dupUA := range s1List {
		cnt += dupUA.dup
		d, _ := Detect(dupUA.ua, true)
		if d == nil || d.DetectedInfo.Model == "" {
			errCnt += dupUA.dup
		}
	}
	fmt.Println("Detect: ", 1-float32(errCnt)/float32(cnt))
}
