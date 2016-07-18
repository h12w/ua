package ua

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	for _, testcase := range []struct {
		ua string
		p  ParsedInfo
	}{
		{
			`Dalvik/1.2.0 (Linux; U; Android 2.1; TR718D Build/MASTER)`,
			ParsedInfo{
				Android:  Version{"2.1", 2, 1, 0},
				Dalvik:   "1.2.0",
				Linux:    true,
				Security: "U",
				Model:    "TR718D",
				Build:    "MASTER",
				OS: OS{
					Name:    "Android",
					Version: Version{"2.1", 2, 1, 0},
				},
			},
		},
		{
			`Mozilla/5.0 (Linux; Android 4.4.2; zh-hk; 7040N Build/KVT49L) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36`,
			ParsedInfo{
				Mozilla: "5.0",
				Linux:   true,
				Android: Version{"4.4.2", 4, 4, 2},
				Locale:  "zh-hk",
				Model:   "7040N",
				Build:   "KVT49L",
				Webkit:  "537.36",
				Version: "4.0",
				Chrome:  "30.0.0.0",
				Mobile:  true,
				Safari:  "537.36",
				OS: OS{
					Name:    "Android",
					Version: Version{"4.4.2", 4, 4, 2},
				},
			},
		},
		{
			`Opera/9.80 (Android 4.2.1; Linux; Opera Tablet/ADR-1212030829) Presto/2.11.355 Version/12.10`,
			ParsedInfo{
				Opera:   "9.80",
				Android: Version{"4.2.1", 4, 2, 1},
				Linux:   true,
				Model:   "Opera",
				Tablet:  true,
				Build:   "ADR-1212030829",
				Presto:  "2.11.355",
				Version: "12.10",
				OS: OS{
					Name:    "Android",
					Version: Version{"4.2.1", 4, 2, 1},
				},
			},
		},
		{
			`Dalvik/1.2.0 (Android 2.2.3 CPU 800MHz)`,
			ParsedInfo{
				Dalvik:  "1.2.0",
				Android: Version{"2.2.3 CPU 800MHz", 2, 2, 3},
				OS: OS{
					Name:    "Android",
					Version: Version{"2.2.3 CPU 800MHz", 2, 2, 3},
				},
			},
		},
		{
			`Dalvik/1.6.0 (Android 4.1.2; GT-I9100 MIUI/2.10.19)`,
			ParsedInfo{
				Android: Version{"4.1.2", 4, 1, 2},
				Dalvik:  "1.6.0",
				Model:   "GT-I9100",
				MIUI:    "2.10.19",
				OS: OS{
					Name:    "Android",
					Version: Version{"4.1.2", 4, 1, 2},
				},
			},
		},
		{
			`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.111 Safari/537.36`,
			ParsedInfo{
				Mozilla:   "5.0",
				Macintosh: true,
				Model:     "Mac",
				OSX:       Version{"10_10_2", 10, 10, 2},
				Webkit:    "537.36",
				Chrome:    "40.0.2214.111",
				Safari:    "537.36",
				OS: OS{
					Name:    "OSX",
					Version: Version{"10_10_2", 10, 10, 2},
				},
			},
		},
		{
			`Mozilla/5.0 (iPad; CPU OS 3_2 like Mac OS X) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B367 Safari/531.21.10`,
			ParsedInfo{
				Mozilla:   "5.0",
				IPad:      true,
				IOS:       Version{"3_2", 3, 2, 0},
				Model:     "iPad",
				Webkit:    "531.21.10",
				Version:   "4.0.4",
				Mobile:    true,
				MobileVer: "7B367",
				Safari:    "531.21.10",
				OS: OS{
					Name:    "iOS",
					Version: Version{"3_2", 3, 2, 0},
				},
			},
		},
		{
			`Mozilla/5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) CriOS/26.0.1410.53 Mobile/11A465 Safari/8536.25 (698D41E5-6572-466A-A970-40B4C19D020F)`,
			ParsedInfo{
				Mozilla:   "5.0",
				IPad:      true,
				Model:     "iPad",
				IOS:       Version{"7_0", 7, 0, 0},
				Webkit:    "536.26",
				CriOS:     "26.0.1410.53",
				Mobile:    true,
				MobileVer: "11A465",
				Safari:    "8536.25",
				OS: OS{
					Name:    "iOS",
					Version: Version{"7_0", 7, 0, 0},
				},
			},
		},
		{
			`GT-I9300 Linux/3.0.13 Android/4.1.9 Release/01.04.2013 Profile/MIDP-2.0 Configuration/CLDC-1.1 Mobile Safari/534.30`,
			ParsedInfo{
				Model:    "GT-I9300",
				Linux:    true,
				LinuxVer: "3.0.13",
				Android:  Version{"4.1.9", 4, 1, 9},
				Release:  "01.04.2013",
				Profile:  "MIDP-2.0",
				Config:   "CLDC-1.1",
				Mobile:   true,
				Safari:   "534.30",
				OS: OS{
					Name:    "Android",
					Version: Version{"4.1.9", 4, 1, 9},
				},
			},
		},
		{
			`Mozilla/5.0 (Linux; Android 4.4.2; SGH-T399N) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/5`,
			ParsedInfo{
				Mozilla: "5.0",
				Linux:   true,
				Android: Version{"4.4.2", 4, 4, 2},
				Model:   "SGH-T399N",
				Webkit:  "537.36",
				Version: "4.0",
				Chrome:  "30.0.0.0",
				Mobile:  true,
				Safari:  "5",
				OS: OS{
					Name:    "Android",
					Version: Version{"4.4.2", 4, 4, 2},
				},
			},
		},
		{
			`HTC_SensationXE_Beats_Z715e_Mozilla/5.0 (Linux; U; Android 2.3.4;it-it;Build/GRH78C) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16`,
			ParsedInfo{
				Model:    "HTC_SensationXE_Beats_Z715e_Mozilla",
				ModelVer: "5.0",
				Linux:    true,
				Security: "U",
				Android:  Version{"2.3.4", 2, 3, 4},
				Locale:   "it-it",
				Build:    "GRH78C",
				Webkit:   "533.16",
				Version:  "5.0",
				Safari:   "533.16",
				OS: OS{
					Name:    "Android",
					Version: Version{"2.3.4", 2, 3, 4},
				},
			},
		},
	} {
		device, err := Parse(testcase.ua, false)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(device.ParsedInfo, testcase.p) {
			t.Fatalf("fail to parse %s, expected \n%#v, got \n%#v", testcase.ua, testcase.p, device.ParsedInfo)
		}
	}
}
