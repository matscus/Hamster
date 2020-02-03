package handlers

import (
	"encoding/xml"
	"log"
	"math/rand"
	"net/http"

	"github.com/matscus/Hamster/Mock/contract_service/structs"
)

type IntRange struct {
	min, max int
}

func IssueComplete(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
	var res structs.IssueCompleteResponse
	res.SOAPENV = "http://schemas.xmlsoap.org/soap/envelope/"
	res.Body.IssueCompleteResponse.Ns2 = "http://tempuri.org/"
	res.Body.IssueCompleteResponse.Ns3 = "http://schemas.microsoft.com/2003/10/Serialization/"
	res.Body.IssueCompleteResponse.Xmlns = "http://schemas.datacontract.org/2004/07/WcfLoanIssueService"
	res.Body.IssueCompleteResponse.IssueCompleteResult.ResultMessage = ""
	res.Body.IssueCompleteResponse.IssueCompleteResult.ResultCode = 0
	rnd := rand.New(rand.NewSource(55))
	ir := IntRange{00000001, 99999999}
	transInfo := make([]structs.TransferInfo, 1, 1)
	transInfo[0].DocId = ir.nextRandom(rnd)
	transInfo[0].TransferType = 1
	res.Body.IssueCompleteResponse.IssueCompleteResult.TransferData.TransferInfo = transInfo
	resMarhal, err := xml.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	resByte := []byte("<?xml version=\"1.0\" encoding=\"utf-16\"?>")
	resByte = append(resByte, resMarhal...)
	w.Header().Set("content-type ", "text/xml")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resByte)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}
func (ir *IntRange) nextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}
