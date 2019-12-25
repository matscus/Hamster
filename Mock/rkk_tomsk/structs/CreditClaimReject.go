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
	ReturnCode        int      `xml:"ReturnCode"`
	InstitutionId     int      `xml:"InstitutionId"`
	ContractReqId     int      `xml:"ContractReqId"`
	ContractReqStatus string   `xml:"ContractReqStatus"`
	CardReqId         int      `xml:"CardReqId"`
	CardReqStatus     string   `xml:"CardReqStatus"`
	ContractId        int      `xml:"ContractId"`
	ContractStatus    string   `xml:"ContractStatus"`
	WorkDay           string   `xml:"WorkDay"`
}
