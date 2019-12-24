package structs

import "encoding/xml"

type ClientDataUpdate struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Tem     string   `xml:"tem,attr"`
	Wcf     string   `xml:"wcf,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text             string `xml:",chardata"`
		ClientDataUpdate struct {
			Text    string `xml:",chardata"`
			Request struct {
				Text           string `xml:",chardata"`
				Benf           string `xml:"Benf"`
				BranchCode     string `xml:"BranchCode"`
				ChannelRep     string `xml:"ChannelRep"`
				Crs            string `xml:"Crs"`
				CrsCountryCode string `xml:"CrsCountryCode"`
				Email          string `xml:"Email"`
				FactAddress    struct {
					Text         string `xml:",chardata"`
					AddrMode     string `xml:"AddrMode"`
					Address      string `xml:"Address"`
					Area         string `xml:"Area"`
					City         string `xml:"City"`
					Construction string `xml:"Construction"`
					Country      string `xml:"Country"`
					Flat         string `xml:"Flat"`
					Frame        string `xml:"Frame"`
					House        string `xml:"House"`
					PostIndex    string `xml:"PostIndex"`
					Region       string `xml:"Region"`
					Street       string `xml:"Street"`
					Town         string `xml:"Town"`
				} `xml:"FactAddress"`
				Fatca         string `xml:"Fatca"`
				InstitutionId string `xml:"InstitutionId"`
				Itin          string `xml:"Itin"`
				Pdl           string `xml:"Pdl"`
				RealUser      string `xml:"RealUser"`
				RegAddress    struct {
					Text         string `xml:",chardata"`
					AddrMode     string `xml:"AddrMode"`
					Address      string `xml:"Address"`
					Area         string `xml:"Area"`
					City         string `xml:"City"`
					Construction string `xml:"Construction"`
					Country      string `xml:"Country"`
					Flat         string `xml:"Flat"`
					Frame        string `xml:"Frame"`
					House        string `xml:"House"`
					PostIndex    string `xml:"PostIndex"`
					Region       string `xml:"Region"`
					Street       string `xml:"Street"`
					Town         string `xml:"Town"`
				} `xml:"RegAddress"`
				Requester string `xml:"Requester"`
				Ssn       string `xml:"Ssn"`
			} `xml:"request"`
		} `xml:"ClientDataUpdate"`
	} `xml:"Body"`
}
