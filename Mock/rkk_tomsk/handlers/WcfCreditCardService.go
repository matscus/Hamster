package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func WcfCreditCardService(w http.ResponseWriter, r *http.Request) {
	log.Println("method ", r.Method)
	for k, v := range r.Header {
		log.Printf("Header %s values %s", k, v)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	log.Println("body ", string(body))
	w.Header().Set("content-type ", "text/xml")
	w.WriteHeader(http.StatusOK)
}
