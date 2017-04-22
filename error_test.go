package ua

import (
	"testing"
)

func TestErrorCases(t *testing.T) {
	ua := "browser: Chrome 32	operating system: Android 4.1	primarily used on: mobile"
	if _, err := Detect(ua, false); err != nil {
		t.Fatal(err)
	}
}
