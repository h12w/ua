package ua

func (d *Device) OSName() string {
	if d.ParsedInfo.OS.Name != "" {
		return d.ParsedInfo.OS.Name
	}
	return d.DetectedInfo.OS.Name
}

func (d *Device) OSVer() string {
	if d.ParsedInfo.OS.Ver != "" {
		return d.ParsedInfo.OS.Ver
	}
	return d.DetectedInfo.OS.Ver
}
