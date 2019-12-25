package structs

import "encoding/xml"

type ClaimStatusRequest struct {
	XMLName        xml.Name `xml:"ClaimStatusRequest"`
	Text           string   `xml:",chardata"`
	Xsd            string   `xml:"xsd,attr"`
	Xsi            string   `xml:"xsi,attr"`
	BranchCode     string   `xml:"BranchCode"`
	Requester      string   `xml:"Requester"`
	InitiatorReqId string   `xml:"InitiatorReqId"`
	DateRequest    string   `xml:"DateRequest"`
}

type ClaimStatusResponse struct {
	XMLName            xml.Name `xml:"ClaimStatusResponse"`
	Text               string   `xml:",chardata"`
	Xsd                string   `xml:"xsd,attr"`
	Xsi                string   `xml:"xsi,attr"`
	ReturnCode         int      `xml:"ReturnCode"`
	InstitutionId      int      `xml:"InstitutionId"`
	ContractReqId      int      `xml:"ContractReqId"`
	ContractReqStatus  string   `xml:"ContractReqStatus"`
	CardReqId          int      `xml:"CardReqId"`
	CardReqStatus      string   `xml:"CardReqStatus"`
	CardBlankReqStatus string   `xml:"CardBlankReqStatus"`
	ContractId         int      `xml:"ContractId"`
	ContractStatus     string   `xml:"ContractStatus"`
	CardReqLimitStatus int      `xml:"CardReqLimitStatus"`
	CardNumberOut      int      `xml:"CardNumberOut"`
	DateExpirationOut  string   `xml:"DateExpirationOut"`
	CardEmbossingOut   string   `xml:"CardEmbossingOut"`
	CrPsk              float64  `xml:"CrPsk"`
	CrPskMoney         float64  `xml:"CrPskMoney"`
	WorkDay            string   `xml:"WorkDay"`
}
