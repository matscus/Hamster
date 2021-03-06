package jmxparser

type TransactionController struct {
	Text      string `xml:",chardata"`
	Guiclass  string `xml:"guiclass,attr"`
	Testclass string `xml:"testclass,attr"`
	Testname  string `xml:"testname,attr"`
	Enabled   string `xml:"enabled,attr"`
	BoolProp  []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"boolProp"`
} //`xml:"TransactionController"`
