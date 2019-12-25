package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func CreditCardContractsHandler(w http.ResponseWriter, r *http.Request) {
	var res structs.CreditContractResponse
	res.ReturnCode = 11
	res.Message = "ФИО и У/Л не найдены"
	w.WriteHeader(http.StatusOK)
	err := xml.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}
