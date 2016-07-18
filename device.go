package ua

type Device struct {
	UA             string
	ParsedInfo     `json:"parsed_info"`
	DetectedInfo   `json:"detected_info"`
	LenientParsing bool
}

type ParsedInfo struct {
	Android   Version
	Build     string
	Chrome    string
	Config    string
	CriOS     string
	Dalvik    string
	IOS       Version
	IPad      bool
	IPhone    bool
	IPod      bool
	Linux     bool
	LinuxVer  string
	Locale    string
	Macintosh bool
	MIUI      string
	Mobile    bool
	MobileVer string
	Model     string
	ModelVer  string
	Mozilla   string
	Opera     string
	OS        OS
	OSX       Version
	Presto    string
	Profile   string
	Release   string
	Safari    string
	Security  string
	Tablet    bool
	Version   string
	Webkit    string
}

type DetectedInfo struct {
	Brand      string
	Model      string
	OS         OS
	IsTablet   bool
	IsWireless bool
	Screen     Size
	Resolution Size
}

type OS struct {
	Name string
	Version
}

type Version struct {
	Ver   string
	Major int
	Minor int
	Patch int
}

type Size struct {
	Width  int
	Height int
}
