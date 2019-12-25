package structs

import "encoding/xml"

type ClientUpdateRequest struct {
	XMLName       xml.Name `xml:"ClientUpdateRequest"`
	Text          string   `xml:",chardata"`
	Xsd           string   `xml:"xsd,attr"`
	Xsi           string   `xml:"xsi,attr"`
	BranchCode    string   `xml:"BranchCode"`
	Requester     string   `xml:"Requester"`
	InstitutionId string   `xml:"InstitutionId"`
	RegAddress    struct {
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
	} `xml:"FactAddress"`
	ChannelRep     string `xml:"ChannelRep"`
	Pdl            string `xml:"Pdl"`
	Fatca          string `xml:"Fatca"`
	Benf           string `xml:"Benf"`
	RealUser       string `xml:"RealUser"`
	Crs            string `xml:"Crs"`
	CrsCountryCode string `xml:"CrsCountryCode"`
}

type ClientUpdateResponse struct {
	XMLName    xml.Name `xml:"ClientUpdateResponse"`
	Text       string   `xml:",chardata"`
	Xsd        string   `xml:"xsd,attr"`
	Xsi        string   `xml:"xsi,attr"`
	ReturnCode string   `xml:"ReturnCode"`
}
