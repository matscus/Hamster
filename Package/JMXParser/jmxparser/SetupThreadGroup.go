package jmxparser

type SetupThreadGroup struct {
	Text       string `xml:",chardata"`
	Guiclass   string `xml:"guiclass,attr"`
	Testclass  string `xml:"testclass,attr"`
	Testname   string `xml:"testname,attr"`
	Enabled    string `xml:"enabled,attr"`
	StringProp []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"stringProp"`
	ElementProp struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		ElementType string `xml:"elementType,attr"`
		Guiclass    string `xml:"guiclass,attr"`
		Testclass   string `xml:"testclass,attr"`
		Testname    string `xml:"testname,attr"`
		Enabled     string `xml:"enabled,attr"`
		BoolProp    struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"boolProp"`
		StringProp struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"stringProp"`
	} `xml:"elementProp"`
	BoolProp struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"boolProp"`
}
