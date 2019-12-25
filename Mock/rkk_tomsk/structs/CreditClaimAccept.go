package structs

import "encoding/xml"

type CreditClaimAcceptRequest struct {
	XMLName        xml.Name `xml:"CreditClaimAcceptRequest"`
	Text           string   `xml:",chardata"`
	Xsd            string   `xml:"xsd,attr"`
	Xsi            string   `xml:"xsi,attr"`
	BranchCode     string   `xml:"BranchCode"`
	Requester      string   `xml:"Requester"`
	InitiatorReqId string   `xml:"InitiatorReqId"`
	DateRequest    string   `xml:"DateRequest"`
	ContractNum    string   `xml:"ContractNum"`
	ControlInfo    string   `xml:"ControlInfo"`
}
type CreditClaimAcceptResponse struct {
	XMLName        xml.Name `xml:"CreditClaimAcceptResponse"`
	Text           string   `xml:",chardata"`
	Xsd            string   `xml:"xsd,attr"`
	Xsi            string   `xml:"xsi,attr"`
	ReturnCode     int      `xml:"ReturnCode"`
	InstitutionId  int      `xml:"InstitutionId"`
	ContractReqId  int      `xml:"ContractReqId"`
	CardReqId      int      `xml:"CardReqId"`
	CardReqStatus  string   `xml:"CardReqStatus"`
	ContractId     int      `xml:"ContractId"`
	ContractStatus string   `xml:"ContractStatus"`
	ResourceId     int      `xml:"ResourceId"`
	AccNumber      string   `xml:"AccNumber"`
	WorkDay        string   `xml:"WorkDay"`
}
