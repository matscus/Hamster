package jmxparser

//TestAction - struct for Test action (pacing)
type TestAction struct {
	Text      string `xml:",chardata"`
	Guiclass  string `xml:"guiclass,attr"`
	Testclass string `xml:"testclass,attr"`
	Testname  string `xml:"testname,attr"`
	Enabled   string `xml:"enabled,attr"`
	IntProp   []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"intProp"`
	StringProp struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"stringProp"`
}
