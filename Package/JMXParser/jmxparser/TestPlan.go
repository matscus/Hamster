package jmxparser

type TestPlan struct {
	Text       string `xml:",chardata"`
	Guiclass   string `xml:"guiclass,attr"`
	Testclass  string `xml:"testclass,attr"`
	Testname   string `xml:"testname,attr"`
	Enabled    string `xml:"enabled,attr"`
	StringProp []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"stringProp"`
	BoolProp []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"boolProp"`
	ElementProp struct {
		Text           string `xml:",chardata"`
		Name           string `xml:"name,attr"`
		ElementType    string `xml:"elementType,attr"`
		Guiclass       string `xml:"guiclass,attr"`
		Testclass      string `xml:"testclass,attr"`
		Testname       string `xml:"testname,attr"`
		Enabled        string `xml:"enabled,attr"`
		CollectionProp struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"collectionProp"`
	} `xml:"elementProp"`
}
