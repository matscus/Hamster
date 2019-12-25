package structs

import "encoding/xml"

type CreditClaimRejectRequest struct {
	XMLName        xml.Name `xml:"CreditClaimRejectRequest"`
	Text           string   `xml:",chardata"`
	Xsd            string   `xml:"xsd,attr"`
	Xsi            string   `xml:"xsi,attr"`
	BranchCode     string   `xml:"BranchCode"`
	Requester      string   `xml:"Requester"`
	InitiatorReqId string   `xml:"InitiatorReqId"`
	DateRequest    string   `xml:"DateRequest"`
	RkkCommReject  string   `xml:"RkkCommReject"`
}
type CreditClaimRejectResponse struct {
	XMLName           xml.Name `xml:"CreditClaimRejectResponse"`
	Text              string   `xml:",chardata"`
	Xsd               string   `xml:"xsd,attr"`
	Xsi               string   `xml:"xsi,attr"`
	ReturnCode        string   `xml:"ReturnCode"`
	InstitutionId     string   `xml:"InstitutionId"`
	ContractReqId     string   `xml:"ContractReqId"`
	ContractReqStatus string   `xml:"ContractReqStatus"`
	CardReqId         string   `xml:"CardReqId"`
	CardReqStatus     string   `xml:"CardReqStatus"`
	ContractId        string   `xml:"ContractId"`
	ContractStatus    string   `xml:"ContractStatus"`
	WorkDay           string   `xml:"WorkDay"`
}
