package handlers

import "net/http"

func WcfCreditCardService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type ", "text/xml")
	w.WriteHeader(http.StatusOK)
}
