package ua

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Parse(ua string, lenient bool) (*Device, error) {
	d := Device{LenientParsing: lenient}
	return d.parse(ua)
}

func (d *Device) parse(ua string) (*Device, error) {
	d.UA = ua
	products, err := scan(ua)
	if err != nil {
		return nil, err
	}
	if err := d.parseFirst(&products[0]); err != nil && !d.LenientParsing {
		return nil, err
	}
	products = products[1:]
	for i := range products {
		if err := d.parseRest(&products[i]); err != nil && !d.LenientParsing {
			return nil, err
		}
	}
	switch d.ParsedInfo.OS.Name {
	case "Android":
		d.ParsedInfo.OS.Version = d.Android
	case "iOS":
		d.ParsedInfo.OS.Version = d.IOS
	case "OSX":
		d.ParsedInfo.OS.Version = d.OSX
	}
	return d, nil
}

func (d *Device) parseFirst(p *product) error {
	switch p.Name {
	case "Mozilla":
		d.Mozilla = p.Version
		return d.parseFirstComments(p.Comments)
	case "Dalvik":
		d.Dalvik = p.Version
		return d.parseFirstComments(p.Comments)
	case "Opera":
		d.Opera = p.Version
		return d.parseFirstComments(p.Comments)
	default:
		d.setModel(p.Name)
		d.ModelVer = p.Version
		return d.parseFirstComments(p.Comments)
	}
	return fmt.Errorf("fail to parse first product: %s", d.UA)
}

func (d *Device) parseRest(p *product) error {
	switch p.Name {
	case "AppleWebKit":
		d.Webkit = p.Version
		return nil
	case "Version":
		d.Version = p.Version
		return nil
	case "Chrome":
		d.Chrome = p.Version
		return nil
	case "Safari":
		d.Safari = p.Version
		return nil
	case "Mobile":
		d.Mobile = true
		d.MobileVer = p.Version
		return nil
	case "CriOS":
		d.CriOS = p.Version
		return nil
	case "Presto":
		d.Presto = p.Version
		return nil
	case "Linux":
		d.Linux = true
		d.LinuxVer = p.Version
		return nil
	case "Android":
		d.ParsedInfo.OS.Name = "Android"
		d.Android = ParseVersion(p.Version)
		return nil
	case "Release":
		d.Release = p.Version
		return nil
	case "Profile":
		d.Profile = p.Version
		return nil
	case "Configuration":
		d.Config = p.Version
		return nil
	}
	return fmt.Errorf("fail to parse product %v: %s", p, d.UA)
}

func (d *Device) parseFirstComments(comments []string) error {
	for i, comment := range comments {
		err := d.parseFirstComment(comment)
		if err != nil {
			if i == len(comments)-1 {
				d.setModel(comment)
				return nil
			}
			return err
		}
	}
	return nil
}

func (d *Device) parseFirstComment(c string) error {
	switch c {
	case "Linux":
		d.Linux = true
		return nil
	case "U", "N", "I":
		d.Security = c
		return nil
	case "iPhone":
		d.ParsedInfo.OS.Name = "iOS"
		d.IPhone = true
		return nil
	case "iPad":
		d.ParsedInfo.OS.Name = "iOS"
		d.IPad = true
		return nil
	case "iPod", "iPod touch":
		d.ParsedInfo.OS.Name = "iOS"
		d.IPod = true
		return nil
	case "Macintosh":
		d.ParsedInfo.OS.Name = "OSX"
		d.Macintosh = true
		return nil
	}
	if err := d.parseLocale(c); err == nil {
		return nil
	}
	tok := strings.Split(c, " ")
	if len(tok) == 0 {
		return nil
	}
	if err := d.parseAndroid(tok); err == nil {
		return nil
	}
	return d.parseDeviceModel(tok)
}

func (d *Device) parseAndroid(tok []string) error {
	if tok[0] == "Android" && len(tok) > 1 {
		d.ParsedInfo.OS.Name = "Android"
		d.Android = ParseVersion(tok[1])
		d.Android.Ver = strings.Join(tok[1:], " ")
		return nil
	}
	return errors.New("not Android comment")
}

var localeRegex = regexp.MustCompile(`(?:^[a-z][a-z][\-_][a-z][a-z]$)|(?:^[a-z][a-z]-$)|(?:^-[a-z][a-z]$)|(?:^-$)`)

func (d *Device) parseLocale(c string) error {
	if localeRegex.MatchString(strings.ToLower(c)) {
		d.Locale = c
		return nil
	}
	return errors.New("not locale")
}

func (d *Device) parseDeviceModel(tok []string) error {
	switch {
	case d.Android.Ver != "":
		return d.parseAndroidModel(tok)
	case d.IPad || d.IPod || d.IPhone:
		return d.parseIOSModel(tok)
	case d.Macintosh:
		return d.parseMacModel(tok)
	}
	return fmt.Errorf("unkown device type: %s", d.UA)
}

func (d *Device) setModel(m string) {
	if d.ParsedInfo.Model == "" {
		d.ParsedInfo.Model = m
	}
}

func (d *Device) parseIOSModel(tok []string) error {
	for i := 0; i <= len(tok)-2; i++ {
		if tok[i] == "OS" {
			switch {
			case d.IPad:
				d.setModel("iPad")
			case d.IPod:
				d.setModel("iPod")
			case d.IPhone:
				d.setModel("iPhone")
			}
			d.IOS = ParseVersion(tok[i+1])
			return nil
		}
	}
	return fmt.Errorf("fail to parse IOS model %v: %s", tok, d.UA)
}

func (d *Device) parseMacModel(tok []string) error {
	for i := 0; i <= len(tok)-2; i++ {
		if tok[i] == "X" {
			d.OSX = ParseVersion(strings.Join(tok[i+1:], " "))
			d.setModel("Mac")
			return nil
		}
	}
	return fmt.Errorf("fail to parse Mac model %v: %s", tok, d.UA)
}

func (d *Device) parseAndroidModel(tok []string) error {
	buildPos, err := d.parseVer(tok)
	if err != nil {
		return err
	}
	d.setModel(strings.Join(tok[:buildPos], " "))
	return nil
}
func (d *Device) parseVer(tok []string) (int, error) {
	for i := 0; i < len(tok); i++ {
		if strings.HasPrefix(tok[i], "Build/") {
			d.Build = strings.TrimPrefix(tok[i], "Build/") + strings.Join(tok[i+1:], " ")
			return i, nil
		} else if strings.HasPrefix(tok[i], "MIUI/") {
			d.MIUI = strings.TrimPrefix(tok[i], "MIUI/") + strings.Join(tok[i+1:], " ")
			return i, nil
		} else if strings.HasPrefix(tok[i], "Tablet/") {
			d.Tablet = true
			d.Build = strings.TrimPrefix(tok[i], "Tablet/") + strings.Join(tok[i+1:], " ")
			return i, nil
		}
	}
	return 0, fmt.Errorf("no 'Build' or 'MIUI' found in: %v, UA: %s", tok, d.UA)
}

func ParseVersion(ver string) (v Version) {
	v.Ver = ver
	vs := strings.Split(ver, ".")
	if len(vs) == 1 {
		vs = strings.Split(ver, "_")
	}
	if len(vs) >= 1 {
		v.Major = atoi(vs[0])
	}
	if len(vs) >= 2 {
		v.Minor = atoi(vs[1])
	}
	if len(vs) >= 3 {
		v.Patch = atoi(vs[2])
	}
	return
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
