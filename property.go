package ua

import "fmt"

func (v *Version) String() string {
	if v.Major != 0 || v.Minor != 0 || v.Patch != 0 {
		return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	}
	return v.Ver
}

func (d *Device) OSName() string {
	if d.ParsedInfo.OS.Name != "" {
		return d.ParsedInfo.OS.Name
	}
	return d.DetectedInfo.OS.Name
}

func (d *Device) OSVer() string {
	if d.ParsedInfo.OS.Version.String() != "" {
		return d.ParsedInfo.OS.Version.String()
	}
	return d.DetectedInfo.OS.Version.String()
}
