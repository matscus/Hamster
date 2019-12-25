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
type CreditContractResponse struct {
	XMLName    xml.Name `xml:"CreditContractResponse"`
	Text       string   `xml:",chardata"`
	Xsd        string   `xml:"xsd,attr"`
	Xsi        string   `xml:"xsi,attr"`
	ReturnCode int      `xml:"ReturnCode"`
	Message    string   `xml:"Message"`
}
