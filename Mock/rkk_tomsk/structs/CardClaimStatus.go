package structs

import "encoding/xml"

type CardClaimStatus struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Tem     string   `xml:"tem,attr"`
	Wcf     string   `xml:"wcf,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text            string `xml:",chardata"`
		CardClaimStatus struct {
			Text    string `xml:",chardata"`
			Request struct {
				Text           string `xml:",chardata"`
				BranchCode     string `xml:"BranchCode"`
				DateRequest    string `xml:"DateRequest"`
				InitiatorReqId string `xml:"InitiatorReqId"`
				Requester      string `xml:"Requester"`
			} `xml:"request"`
		} `xml:"CardClaimStatus"`
	} `xml:"Body"`
}
