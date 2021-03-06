package jmxparser

//ConstantThroughputTimer - func for Constant Throughput Timer
type ConstantThroughputTimer struct {
	Text      string `xml:",chardata"`
	Guiclass  string `xml:"guiclass,attr"`
	Testclass string `xml:"testclass,attr"`
	Testname  string `xml:"testname,attr"`
	Enabled   string `xml:"enabled,attr"`
	IntProp   struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"intProp"`
	StringProp struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"stringProp"`
}
