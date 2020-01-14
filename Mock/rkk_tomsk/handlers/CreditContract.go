package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func CreditCardContractsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("body 1 ", r.Body)
	var res structs.CreditCardContractsResponse
	res.SOAPENV = "http://schemas.xmlsoap.org/soap/envelope/"
	res.Body.CreditCardContractsResponse.CreditCardContractsResult.CredInfoList.Nil = true
	res.Body.CreditCardContractsResponse.CreditCardContractsResult.CredInfoList.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	res.Body.CreditCardContractsResponse.CreditCardContractsResult.ReturnCode = 1012
	res.Body.CreditCardContractsResponse.CreditCardContractsResult.Message = "ФИО и У/Л не найдены"
	res.Body.CreditCardContractsResponse.Xmlns = "http://schemas.datacontract.org/2004/07/WcfCreditCardService"
	res.Body.CreditCardContractsResponse.Ns2 = "http://tempuri.org/"
	res.Body.CreditCardContractsResponse.Ns3 = "http://schemas.microsoft.com/2003/10/Serialization/"
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

// `<?xml version="1.0" encoding="utf-16"?>
// 	<SOAP-ENV:Envelope
// 	   xmlns:SOAP-ENV = "http://schemas.xmlsoap.org/soap/envelope/">
// 	   <SOAP-ENV:Body xmlns:m = "http://www.xyz.org/quotation">
// 			<CreditCardContractsResponse xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
// 					<ReturnCode>1012</ReturnCode>
// 					<Message>ФИО и У/Л не найдены</Message>
// 			</CreditContractResponse>
// 	   </SOAP-ENV:Body>
// </SOAP-ENV:Envelope>`
