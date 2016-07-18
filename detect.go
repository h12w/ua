package ua

func Detect(ua string, lenientParsing bool) (*Device, error) {
	d := Device{LenientParsing: lenientParsing}
	_, err := d.parse(ua)
	if err != nil && !lenientParsing {
		return nil, err
	}
	d.detect()
	return &d, nil
}

func (d *Device) detect() {
	if di, ok := deviceInfoMap[d.ParsedInfo.Model]; ok {
		d.DetectedInfo = di
	}
}
