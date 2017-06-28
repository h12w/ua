package ua

import "testing"

func TestModel(t *testing.T) {
	for _, testcase := range []struct {
		model string
		ua    string
	}{
		{"vivo X7", "Mozilla/5.0 (Linux; Android 5.1.1; vivo X7 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/39.0.0.0 Mobile Safari/537.36"},
	} {
		device, err := Parse(testcase.ua, false)
		if err != nil {
			t.Fatal(err)
		}
		if device.ParsedInfo.Model != testcase.model {
			t.Fatalf("fail to parse %s, expected \n%#v, got \n%#v", testcase.ua, testcase.model, device.ParsedInfo.Model)
		}
	}
}
