package ua

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkDetect(b *testing.B) {
	f, err := os.Open("testdata/s1")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	var UAs []string
	for s.Scan() {
		dupUA := strings.SplitN(s.Text(), " ", 2)
		UAs = append(UAs, dupUA[1])
	}
	f.Close()
	if s.Err() != nil {
		panic(s.Err())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ua := range UAs {

			if _, err := Detect(ua, true); err != nil {
				//fmt.Println(err)
			}
		}
	}
	fmt.Printf("%d UAs parsed per loop.\n", len(UAs))
}
