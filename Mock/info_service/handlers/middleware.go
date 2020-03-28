package handlers

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	Mean       float64
	Deviation  float64
	Requestlog bool
)

func Middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if Requestlog {
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
		}
		requestId := r.Header.Get("GPB-requestId")
		w.Header().Set("GPB-requestId", requestId)
		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		if Mean != 0.0 {
			waitResponse()
		}
		f(w, r)
	}
}

func waitResponse() {
	timeout := rand.NormFloat64()*Deviation + Mean
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}
