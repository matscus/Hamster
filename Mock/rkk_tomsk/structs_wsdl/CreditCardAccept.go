package structs

import "encoding/xml"

type CreditCardAccept struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Tem     string   `xml:"tem,attr"`
	Wcf     string   `xml:"wcf,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text             string `xml:",chardata"`
		CreditCardAccept struct {
			Text    string `xml:",chardata"`
			Request struct {
				Text           string `xml:",chardata"`
				BranchCode     string `xml:"BranchCode"`
				ContractNum    string `xml:"ContractNum"`
				ControlInfo    string `xml:"ControlInfo"`
				DateRequest    string `xml:"DateRequest"`
				InitiatorReqId string `xml:"InitiatorReqId"`
				Requester      string `xml:"Requester"`
			} `xml:"request"`
		} `xml:"CreditCardAccept"`
	} `xml:"Body"`
}
