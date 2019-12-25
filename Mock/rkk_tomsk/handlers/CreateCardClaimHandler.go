package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func CreateCardClaimRHandler(w http.ResponseWriter, r *http.Request) {
	var statusReq structs.ClaimStatusRequest
	err := xml.NewDecoder(r.Body).Decode(&statusReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	var statusRes structs.CreateCardClaimResponse
	statusRes.ReturnCode = 11
	statusRes.Message = "У клиента есть действующий договор на кредитную карту"
	statusRes.InstitutionId = 0
	statusRes.ContractReqId = 0
	statusRes.CardReqId = 0
	statusRes.DateExpirationOut = "0001-01-01T00:00:00"
	statusRes.CrPsk = 0
	statusRes.CrPskMoney = 0
	statusRes.WorkDay = "1900-01-01T00:00:00"
}
