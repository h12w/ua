package ua

var deviceInfoMapExt = map[string]DetectedInfo{
	"iris354":  DetectedInfo{Brand: "Lava", Model: "Iris 354", OS: OS{"Android", Version{"4.2", 4, 2, 0}}, IsTablet: false, IsWireless: true, Screen: Size{Width: 50, Height: 74}, Resolution: Size{Width: 320, Height: 480}},
	"SM-G900H": DetectedInfo{Brand: "Samsung", Model: "SM-G900H", OS: OS{"Android", Version{"4.4", 4, 4, 0}}, IsTablet: false, IsWireless: true, Screen: Size{Width: 64, Height: 113}, Resolution: Size{Width: 1080, Height: 1920}},
	"ZTE V797": DetectedInfo{Brand: "ZTE", Model: "ZTE V797", OS: OS{"Android", Version{"4.2", 4, 2, 2}}, IsTablet: false, IsWireless: true, Screen: Size{Width: 34, Height: 50}, Resolution: Size{Width: 320, Height: 480}},
}

func init() {
	for k, v := range deviceInfoMapExt {
		deviceInfoMap[k] = v
	}
}
