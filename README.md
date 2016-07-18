ua: UA parser in Go
===================

Design
------

The UA detection algorithm is a pipeline containing 3 steps:

1. Scan: scan products and its corresponding comments from a UA string.
2. Parse: parse device info from products.
3. Detect: detect device model by looking up in a database and get related info, e.g. screen resolution.

### Scan

The scanner scans the user agent string into products. Each product has its own name, version and comments. e.g.

For user agent string:

```
Mozilla/5.0 (Linux; U; Android ROM v3; en-us; ALCATEL ONE TOUCH 991 Build/GRK39F) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1
```

Its products are:

* Mozilla/5.0
	- Linux
	- U
	- Android ROM v3
	- en-us
	- ALCATEL ONE TOUCH 991 Build/GRK39F
* AppleWebKit/533.1 
	- KHTML, like Gecko
* Version/4.0
* Mobile
* Safari/533.1

### Parse

The parser then parses the scanned products and fill the information into a Device struct, e.g.

```
Mozilla:  5.0
Linux:    true
Security: U
Android:  ROM v3
Locale:   en-us
Model:    ALCATEL ONE TOUCH 991
Build:    GRK39F
Webkit:   533.1
Version:  4.0
Mobile:   true
Safari:   533.1
```

### Detect

The detector is then able to use the parsed Device structure to detect the device type & related info: vendor, screen resolution, tablet-or-not, etc, by searching in a device database.

Design Goals
------------

### Accuracy

The accuracy of the detection is determined by the completeness of the device database, then we might have to migrate the data from multiple sources. Automatic tests will be adapted from those resources.

References
----------

### Detectors
* [WURFL](wurfl.sourceforge.net)
* [HandsetDetection](http://www.handsetdetection.com)
* [DeviceMap](https://devicemap.apache.org)
* [Mobile-Detect](https://github.com/serbanghita/Mobile-Detect)
* [Device-Detector](https://github.com/piwik/device-detector)
* [Detector](https://github.com/dmolsen/Detector).

### Parsers
* [user agent](https://github.com/mssola/user_agent)
* [uaparser](https://github.com/varstr/uaparser)
