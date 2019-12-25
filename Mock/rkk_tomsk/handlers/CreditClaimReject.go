package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func CreditClaimRejectHandler(w http.ResponseWriter, r *http.Request) {
	var req structs.ClaimStatusRequest
	err := xml.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	//добавить обработку если будет нужно
	var res structs.CreditClaimRejectResponse
	res.ReturnCode = 0
	res.InstitutionId = 588711
	res.ContractReqId = 439657
	res.ContractReqStatus = "На ожидании подписи клиента"
	res.CardReqId = 1091099
	res.CardReqStatus = "На ожидании выдачи"
	res.ContractId = 1058693
	res.ContractStatus = "Введен"
	res.WorkDay = "1900-01-01T00:00:00"
}
