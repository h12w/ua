package ua

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	for _, testcase := range []struct {
		ua       string
		products []product
	}{
		{
			`Mozilla/5.0 (Linux; U; Android RUS MOD V2.991; ru-ru; ALCATEL ONE TOUCH 991 Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
			[]product{
				{"Mozilla", "5.0", []string{"Linux", "U", "Android RUS MOD V2.991", "ru-ru", "ALCATEL ONE TOUCH 991 Build/GRJ90"}},
				{"AppleWebKit", "533.1", []string{"KHTML, like Gecko"}},
				{"Version", "4.0", nil},
				{"Mobile", "", nil},
				{"Safari", "533.1", nil},
			},
		},
		{
			`Dalvik/1.6.0 (Linux; U; Android L (5.0.2); teXet_X-medium_plus Build/JDQ39)`,
			[]product{
				{"Dalvik", "1.6.0", []string{"Linux", "U", "Android L (5.0.2)", "teXet_X-medium_plus Build/JDQ39"}},
			},
		},
		{
			`Mobile/11A465  (3B92C18B-D9DE-4CB7-A02A-22FD2AF17C8F)`,
			[]product{
				{"Mobile", "11A465", []string{"3B92C18B-D9DE-4CB7-A02A-22FD2AF17C8F"}},
			},
		},
		{
			`Mozilla/5.0(Linux; U)`,
			[]product{
				{"Mozilla", "5.0", []string{"Linux", "U"}},
			},
		},
	} {
		products, err := scan(testcase.ua)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(products, testcase.products) {
			t.Fatalf("fail to scan %s --- %#v", testcase.ua, products)
		}
	}
}
