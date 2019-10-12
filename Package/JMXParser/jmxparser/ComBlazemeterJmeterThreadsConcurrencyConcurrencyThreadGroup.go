package jmxparser

type ComBlazemeterJmeterThreadsConcurrencyConcurrencyThreadGroup struct {
	Text        string `xml:",chardata"`
	Guiclass    string `xml:"guiclass,attr"`
	Testclass   string `xml:"testclass,attr"`
	Testname    string `xml:"testname,attr"`
	Enabled     string `xml:"enabled,attr"`
	ElementProp struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		ElementType string `xml:"elementType,attr"`
	} `xml:"elementProp"`
	StringProp []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"stringProp"`
}
