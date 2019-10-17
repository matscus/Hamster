package jmxparser

//ResultCollector - struct for result collector
type ResultCollector struct {
	Text      string `xml:",chardata"`
	Guiclass  string `xml:"guiclass,attr"`
	Testclass string `xml:"testclass,attr"`
	Testname  string `xml:"testname,attr"`
	Enabled   string `xml:"enabled,attr"`
	BoolProp  struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"boolProp"`
	ObjProp struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name"`
		Value struct {
			Text                               string `xml:",chardata"`
			Class                              string `xml:"class,attr"`
			Time                               string `xml:"time"`
			Latency                            string `xml:"latency"`
			Timestamp                          string `xml:"timestamp"`
			Success                            string `xml:"success"`
			Label                              string `xml:"label"`
			Code                               string `xml:"code"`
			Message                            string `xml:"message"`
			ThreadName                         string `xml:"threadName"`
			DataType                           string `xml:"dataType"`
			Encoding                           string `xml:"encoding"`
			Assertions                         string `xml:"assertions"`
			Subresults                         string `xml:"subresults"`
			ResponseData                       string `xml:"responseData"`
			SamplerData                        string `xml:"samplerData"`
			XML                                string `xml:"xml"`
			FieldNames                         string `xml:"fieldNames"`
			ResponseHeaders                    string `xml:"responseHeaders"`
			RequestHeaders                     string `xml:"requestHeaders"`
			ResponseDataOnError                string `xml:"responseDataOnError"`
			SaveAssertionResultsFailureMessage string `xml:"saveAssertionResultsFailureMessage"`
			AssertionsResultsToSave            string `xml:"assertionsResultsToSave"`
			Bytes                              string `xml:"bytes"`
			SentBytes                          string `xml:"sentBytes"`
			URL                                string `xml:"url"`
			ThreadCounts                       string `xml:"threadCounts"`
			IdleTime                           string `xml:"idleTime"`
			ConnectTime                        string `xml:"connectTime"`
		} `xml:"value"`
	} `xml:"objProp"`
	StringProp struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"stringProp"`
}
