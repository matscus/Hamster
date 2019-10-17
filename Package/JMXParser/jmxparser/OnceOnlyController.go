package jmxparser

//OnceOnlyController - struct for Once controller
type OnceOnlyController struct {
	Text      string `xml:",chardata"`
	Guiclass  string `xml:"guiclass,attr"`
	Testclass string `xml:"testclass,attr"`
	Testname  string `xml:"testname,attr"`
	Enabled   string `xml:"enabled,attr"`
}
