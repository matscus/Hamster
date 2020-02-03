package structs

import "encoding/xml"

type IssueComplete struct {
	XMLName xml.Name `xml:"IssueComplete"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Ns2     string   `xml:"ns2,attr"`
	Ns3     string   `xml:"ns3,attr"`
	Request struct {
		Text             string `xml:",chardata"`
		BranchCode       string `xml:"BranchCode"`
		Date             string `xml:"Date"`
		ExtLoanRepayment struct {
			Text        string `xml:",chardata"`
			ExtTransfer struct {
				Text           string `xml:",chardata"`
				Amount         string `xml:"Amount"`
				BankName       string `xml:"BankName"`
				Bic            string `xml:"Bic"`
				Description    string `xml:"Description"`
				ReceiverAcc    string `xml:"ReceiverAcc"`
				ReceiverCorAcc string `xml:"ReceiverCorAcc"`
				ReceiverInn    string `xml:"ReceiverInn"`
				ReceiverKpp    string `xml:"ReceiverKpp"`
				ReceiverName   string `xml:"ReceiverName"`
				SenderAcc      string `xml:"SenderAcc"`
				SenderDul      string `xml:"SenderDul"`
				SenderFio      string `xml:"SenderFio"`
				TransferType   string `xml:"TransferType"`
			} `xml:"ExtTransfer"`
		} `xml:"ExtLoanRepayment"`
		InsuranceTransfer struct {
			Text        string `xml:",chardata"`
			ExtTransfer struct {
				Text           string `xml:",chardata"`
				Amount         string `xml:"Amount"`
				BankName       string `xml:"BankName"`
				Bic            string `xml:"Bic"`
				Description    string `xml:"Description"`
				ReceiverAcc    string `xml:"ReceiverAcc"`
				ReceiverCorAcc string `xml:"ReceiverCorAcc"`
				ReceiverInn    string `xml:"ReceiverInn"`
				ReceiverKpp    string `xml:"ReceiverKpp"`
				ReceiverName   string `xml:"ReceiverName"`
				SenderAcc      string `xml:"SenderAcc"`
				SenderDul      string `xml:"SenderDul"`
				SenderFio      string `xml:"SenderFio"`
				TransferType   string `xml:"TransferType"`
			} `xml:"ExtTransfer"`
		} `xml:"InsuranceTransfer"`
		LoanIssue struct {
			Text         string `xml:",chardata"`
			Amount       string `xml:"Amount"`
			CreditAcc    string `xml:"CreditAcc"`
			DebitAcc     string `xml:"DebitAcc"`
			Description  string `xml:"Description"`
			TransferType string `xml:"TransferType"`
		} `xml:"LoanIssue"`
		Requester string `xml:"Requester"`
	} `xml:"request"`
}
type IssueCompleteResponse struct {
	XMLName xml.Name `xml:"SOAP-ENV:Envelope"`
	Text    string   `xml:",chardata"`
	SOAPENV string   `xml:"xmlns:SOAP-ENV,attr"`
	Header  string   `xml:"SOAP-ENV:Header"`
	Body    struct {
		Text                  string `xml:",chardata"`
		IssueCompleteResponse struct {
			Text                string `xml:",chardata"`
			Xmlns               string `xml:"xmlns,attr"`
			Ns2                 string `xml:"xmlns:ns2,attr"`
			Ns3                 string `xml:"xmlns:ns3,attr"`
			IssueCompleteResult struct {
				Text          string `xml:",chardata"`
				ResultCode    int    `xml:"ResultCode"`
				ResultMessage string `xml:"ResultMessage"`
				TransferData  struct {
					Text         string         `xml:",chardata"`
					TransferInfo []TransferInfo `xml:"TransferInfo"`
				} `xml:"TransferData"`
			} `xml:"ns2:IssueCompleteResult"`
		} `xml:"ns2:IssueCompleteResponse"`
	} `xml:"SOAP-ENV:Body"`
}

type TransferInfo struct {
	Text         string `xml:",chardata"`
	DocId        int    `xml:"DocId"`
	TransferType int    `xml:"TransferType"`
}
