package handlers

import (
	"log"
	"net/http"
)

func WcfCreditCardService(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	w.Header().Set("content-type ", "text/xml")
	w.WriteHeader(http.StatusOK)
}
