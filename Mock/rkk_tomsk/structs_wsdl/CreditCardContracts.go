package structs

import "encoding/xml"

type CreditCardContracts struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Tem     string   `xml:"tem,attr"`
	Wcf     string   `xml:"wcf,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                string `xml:",chardata"`
		CreditCardContracts struct {
			Text    string `xml:",chardata"`
			Request struct {
				Text         string `xml:",chardata"`
				BranchCode   string `xml:"BranchCode"`
				ClientDocNum string `xml:"ClientDocNum"`
				ClientDocSer string `xml:"ClientDocSer"`
				ClientName   string `xml:"ClientName"`
				ClientName1  string `xml:"ClientName1"`
				ClientName2  string `xml:"ClientName2"`
				DateCheck    string `xml:"DateCheck"`
				IsShowClose  string `xml:"IsShowClose"`
				PrPeriod     string `xml:"PrPeriod"`
				Process      string `xml:"Process"`
				Requester    string `xml:"Requester"`
				SelFlag      string `xml:"SelFlag"`
				TargetFund   string `xml:"TargetFund"`
			} `xml:"request"`
		} `xml:"CreditCardContracts"`
	} `xml:"Body"`
}
