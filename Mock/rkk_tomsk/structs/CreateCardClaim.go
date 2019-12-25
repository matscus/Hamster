package structs

import "encoding/xml"

type CreateCardClaimRequest struct {
	XMLName            xml.Name `xml:"CreateCardClaimRequest"`
	Text               string   `xml:",chardata"`
	Xsd                string   `xml:"xsd,attr"`
	Xsi                string   `xml:"xsi,attr"`
	BranchCode         string   `xml:"BranchCode"`
	Requester          string   `xml:"Requester"`
	InitiatorReqId     string   `xml:"InitiatorReqId"`
	InitiatorReqNum    string   `xml:"InitiatorReqNum"`
	DateRequest        string   `xml:"DateRequest"`
	ChannelReq         string   `xml:"ChannelReq"`
	ClientName         string   `xml:"ClientName"`
	ClientName1        string   `xml:"ClientName1"`
	ClientName2        string   `xml:"ClientName2"`
	Sex                string   `xml:"Sex"`
	IsResident         string   `xml:"IsResident"`
	CountryIsoCode     string   `xml:"CountryIsoCode"`
	BirthDate          string   `xml:"BirthDate"`
	BirthPlace         string   `xml:"BirthPlace"`
	ClientDocSMode     string   `xml:"ClientDocSMode"`
	ClientDocTypeCode  string   `xml:"ClientDocTypeCode"`
	ClientDocSer       string   `xml:"ClientDocSer"`
	ClientDocNum       string   `xml:"ClientDocNum"`
	ClientDocIssueDate string   `xml:"ClientDocIssueDate"`
	ClientDocIssuer    string   `xml:"ClientDocIssuer"`
	ClientDocSubCode   string   `xml:"ClientDocSubCode"`
	RegAddress         struct {
		Text      string `xml:",chardata"`
		AddrMode  string `xml:"AddrMode"`
		PostIndex string `xml:"PostIndex"`
		Region    string `xml:"Region"`
		Area      string `xml:"Area"`
		Town      string `xml:"Town"`
		Street    string `xml:"Street"`
		House     string `xml:"House"`
		Flat      string `xml:"Flat"`
	} `xml:"RegAddress"`
	FactAddress struct {
		Text      string `xml:",chardata"`
		AddrMode  string `xml:"AddrMode"`
		PostIndex string `xml:"PostIndex"`
		Region    string `xml:"Region"`
		Area      string `xml:"Area"`
		Town      string `xml:"Town"`
		Street    string `xml:"Street"`
		House     string `xml:"House"`
		Flat      string `xml:"Flat"`
	} `xml:"FactAddress"`
	TempAddress struct {
		Text     string `xml:",chardata"`
		AddrMode string `xml:"AddrMode"`
	} `xml:"TempAddress"`
	TelMob                 string `xml:"TelMob"`
	TelHome                string `xml:"TelHome"`
	CrFund                 string `xml:"CrFund"`
	CrLimit                string `xml:"CrLimit"`
	Pdl                    string `xml:"Pdl"`
	Fatca                  string `xml:"Fatca"`
	Benf                   string `xml:"Benf"`
	RealUser               string `xml:"RealUser"`
	InUsr                  string `xml:"InUsr"`
	NumberDo               string `xml:"NumberDo"`
	Mode                   string `xml:"Mode"`
	DateExpiration         string `xml:"DateExpiration"`
	PLoyalty               string `xml:"PLoyalty"`
	PaymentSystem          string `xml:"PaymentSystem"`
	TypeCard               string `xml:"TypeCard"`
	IsUnembossed           string `xml:"IsUnembossed"`
	CProgram               string `xml:"CProgram"`
	NewInstitutionId       string `xml:"NewInstitutionId"`
	IgnoreCollision        string `xml:"IgnoreCollision"`
	BirthCountryIsoCode    string `xml:"BirthCountryIsoCode"`
	Crs                    string `xml:"Crs"`
	CrsCountryCode         string `xml:"CrsCountryCode"`
	DateIdb                string `xml:"DateIdb"`
	DateBki                string `xml:"DateBki"`
	Idb                    string `xml:"Idb"`
	PaymentAmount          string `xml:"PaymentAmount"`
	RtdMAverageMonthIncome string `xml:"RtdMAverageMonthIncome"`
	ConsLoanDebt           string `xml:"ConsLoanDebt"`
	SeniorityYears         string `xml:"SeniorityYears"`
	SeniorityMonths        string `xml:"SeniorityMonths"`
	CurSeniorityYears      string `xml:"CurSeniorityYears"`
	CurSeniorityMonths     string `xml:"CurSeniorityMonths"`
}

type CreateCardClaimResponse struct {
	XMLName           xml.Name `xml:"CreateCardClaimResponse"`
	Text              string   `xml:",chardata"`
	Xsd               string   `xml:"xsd,attr"`
	Xsi               string   `xml:"xsi,attr"`
	ReturnCode        int      `xml:"ReturnCode"`
	Message           string   `xml:"Message"`
	InstitutionId     int      `xml:"InstitutionId"`
	ContractReqId     int      `xml:"ContractReqId"`
	CardReqId         int      `xml:"CardReqId"`
	DateExpirationOut string   `xml:"DateExpirationOut"`
	CrPsk             int      `xml:"CrPsk"`
	CrPskMoney        int      `xml:"CrPskMoney"`
	WorkDay           string   `xml:"WorkDay"`
}
