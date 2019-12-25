package handlers

import (
	"net/http"
	"reflect"
)

func CreditCardContractsHandler(w http.ResponseWriter, r *http.Request) {
	// var res structs.CreditCardContractsResponse
	// res.XMLName.Space = "\"http://schemas.xmlsoap.org/soap/envelope/\""
	// //res.XMLName.Local = "\"хуёкал\""
	// res.ReturnCode = 11
	// res.Message = "ФИО и У/Л не найдены"
	// w.WriteHeader(http.StatusOK)
	// t, _ := xml.Marshal(res)
	// s := []byte(`<?xml version="1.0" encoding="utf-16"?>
	// <SOAP-ENV:Envelope
	// xmlns:SOAP-ENV = "http://schemas.xmlsoap.org/soap/envelope/">
	// <SOAP-ENV:Body xmlns:m = "http://www.xyz.org/quotation">`)
	// s = append(s, t...)
	// s = append(s, []byte(`</SOAP-ENV:Body>
	// </SOAP-ENV:Envelope>`)...)

	//    </SOAP-ENV:Body>
	// </SOAP-ENV:Envelope>`
	x := []byte(
		`<?xml version="1.0" encoding="utf-16"?>
	<SOAP-ENV:Envelope
	   xmlns:SOAP-ENV = "http://schemas.xmlsoap.org/soap/envelope/">
	   <SOAP-ENV:Body xmlns:m = "http://www.xyz.org/quotation">
			<CreditCardContractsResponse xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
					<ReturnCode>1012</ReturnCode>
					<Message>ФИО и У/Л не найдены</Message>
			</CreditContractResponse>
	   </SOAP-ENV:Body>
	</SOAP-ENV:Envelope>`)
	w.Write(x)

	// err := xml.NewEncoder(w).Encode(x)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	// 	if errWrite != nil {
	// 		log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
	// 	}
	// }
}

func getBytes(msgPart interface{}) (bytes []byte) {
	v := reflect.ValueOf(msgPart)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			continue
		}

		bytes = append(bytes, v.Field(i).Bytes()...)
	}

	return bytes
}
