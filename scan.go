package ua

import (
	"bytes"
	"fmt"
	"strings"
)

type product struct {
	Name     string
	Version  string
	Comments []string
}

func scan(ua string) (products []product, err error) {
	buf := []byte(ua)
	for len(buf) > 0 {
		var p product
		p, buf = scanProduct(buf)
		products = append(products, p)
	}
	if len(products) == 0 {
		err = fmt.Errorf("fail to parse products from: %s", ua)
	}
	return
}

func scanProduct(ua []byte) (s product, _ []byte) {
	token := readToken(ua)
	s.Name, s.Version = scanNameVersion(token)
	i := len(token)
	for i < len(ua) && ua[i] == ' ' {
		i++
	}
	if i < len(ua) && ua[i] == '(' {
		i++
		token := readComments(ua[i:])
		i += len(token) + 1
		s.Comments = splitComments(token)
		i++
	}
	if i > len(ua) {
		i = len(ua)
	}
	return s, ua[i:]
}
func readToken(ua []byte) []byte {
	for i := range ua {
		switch ua[i] {
		case ' ', '\t', '(':
			return ua[:i]
		}
	}
	return ua
}
func readComments(ua []byte) []byte {
	level := 0
	for i := range ua {
		switch ua[i] {
		case ')':
			if level == 0 {
				return ua[:i]
			}
			level--
		case '(':
			level++
		}
	}
	return ua
}
func scanNameVersion(product []byte) (string, string) {
	prod := strings.SplitN(string(product), "/", 2)
	if len(prod) == 2 {
		return prod[0], prod[1]
	}
	return string(product), ""
}
func splitComments(buf []byte) (comments []string) {
	start := 0
	for i := range buf {
		switch buf[i] {
		case ';':
			if comment := bytes.TrimSpace(buf[start:i]); len(comment) != 0 {
				comments = append(comments, string(comment))
			}
			start = i + 1
		}
	}
	if comment := bytes.TrimSpace(buf[start:]); len(comment) != 0 {
		comments = append(comments, string(comment))
	}
	return
}
