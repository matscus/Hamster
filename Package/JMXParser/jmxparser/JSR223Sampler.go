package jmxparser

//JSR223Sampler - struct for JSR223Sampler
type JSR223Sampler []struct {
	Text       string `xml:",chardata"`
	Guiclass   string `xml:"guiclass,attr"`
	Testclass  string `xml:"testclass,attr"`
	Testname   string `xml:"testname,attr"`
	Enabled    string `xml:"enabled,attr"`
	StringProp []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"stringProp"`
}
