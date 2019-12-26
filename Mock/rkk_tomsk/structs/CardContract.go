package structs

import "encoding/xml"

type CreditContractRequest struct {
	XMLName      xml.Name `xml:"CreditContractRequest"`
	Text         string   `xml:",chardata"`
	Xsd          string   `xml:"xsd,attr"`
	Xsi          string   `xml:"xsi,attr"`
	ClientName   string   `xml:"ClientName"`
	ClientName1  string   `xml:"ClientName1"`
	ClientName2  string   `xml:"ClientName2"`
	ClientDocSer string   `xml:"ClientDocSer"`
	ClientDocNum string   `xml:"ClientDocNum"`
	TargetFund   string   `xml:"TargetFund"`
	DateCheck    string   `xml:"DateCheck"`
	SelFlag      string   `xml:"SelFlag"`
	PrPeriod     string   `xml:"PrPeriod"`
	IsShowClose  string   `xml:"IsShowClose"`
	Requester    string   `xml:"Requester"`
	Process      string   `xml:"Process"`
}

// type CreditCardContractsResponse struct {
// 	XMLName    xml.Name `xml:"CreditCardContractsResponse xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"`
// 	Text       string   `xml:",chardata"`
// 	Xsd        string   `xml:"xsd,attr"`
// 	Xsi        string   `xml:"xsi,attr"`
// 	ReturnCode int      `xml:"ReturnCode"`
// 	Message    string   `xml:"Message"`
// }
type CreditCardContractsResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	Text          string   `xml:",chardata"`
	SOAPENV       string   `xml:"SOAP-ENV,attr"`
	EncodingStyle string   `xml:"encodingStyle,attr"`
	Body          struct {
		Text                        string `xml:",chardata"`
		M                           string `xml:"xmlns:m,attr"`
		CreditCardContractsResponse struct {
			XMLName    xml.Name `xml:"CreditCardContractsResponse xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"`
			Text       string   `xml:",chardata"`
			Xsd        string   `xml:"xmlns:xsd,attr"`
			Xsi        string   `xml:"xmlns:xsi,attr"`
			ReturnCode int      `xml:"ReturnCode"`
			Message    string   `xml:"Message"`
		} `xml:"CreditCardContractsResponse"`
	} `xml:"SOAP-ENV:Body"`
}
