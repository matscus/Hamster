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

type CreditCardContractsResponse struct {
	XMLName xml.Name `xml:"SOAP-ENV:Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"xmlns:SOAP-ENV,attr"`
	Header  string   `xml:"SOAP-ENV:Header"`
	Body    struct {
		Text                        string `xml:",chardata"`
		CreditCardContractsResponse struct {
			Text                      string `xml:",chardata"`
			Xmlns                     string `xml:"xmlns,attr"`
			Ns2                       string `xml:"xmlns:ns2,attr"`
			Ns3                       string `xml:"xmlns:ns3,attr"`
			CreditCardContractsResult struct {
				Text         string `xml:",chardata"`
				CredInfoList struct {
					Text string `xml:",chardata"`
					Nil  bool   `xml:"nil,attr"`
					Xsi  string `xml:"xsi,attr"`
				} `xml:"CredInfoList"`
				Message    string `xml:"Message"`
				ReturnCode int    `xml:"ReturnCode"`
			} `xml:"ns2:CreditCardContractsResult"`
		} `xml:"ns2:CreditCardContractsResponse"`
	} `xml:"SOAP-ENV:Body"`
}
