package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/Mock/rkk_tomsk/handlers"
)

var (
	pemPath      string
	keyPath      string
	proto        string
	listenport   string
	wait         time.Duration
	writeTimeout time.Duration
	readTimeout  time.Duration
	idleTimeout  time.Duration
	requestlog   bool
)

func main() {
	flag.StringVar(&pemPath, "pempath", os.Getenv("SERVERREM"), "path to pem file")
	flag.StringVar(&keyPath, "keypath", os.Getenv("SERVERKEY"), "path to key file")
	flag.StringVar(&listenport, "port", "10000", "port to Listen")
	flag.StringVar(&proto, "proto", "http", "http or https")
	flag.BoolVar(&requestlog, "request-log", false, "idle server timeout")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully")
	flag.DurationVar(&readTimeout, "read-timeout", time.Second*15, "read server timeout")
	flag.DurationVar(&writeTimeout, "write-timeout", time.Second*15, "write server timeout")
	flag.DurationVar(&idleTimeout, "idle-timeout", time.Second*60, "idle server timeout")
	flag.Parse()
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         "0.0.0.0:" + listenport,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
	r.HandleFunc("/WcfCreditCardService/CreditCardService.svc", middleware()).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
	http.Handle("/", r)
	r.Use(mux.CORSMethodMiddleware(r))
	go func() {
		switch proto {
		case "https":
			log.Printf("Server is run, proto: https, address: %s ", srv.Addr)
			if err := srv.ListenAndServeTLS(pemPath, keyPath); err != nil {
				log.Println(err)
			}
		case "http":
			log.Printf("Server is run, proto: http, address: %s ", srv.Addr)
			if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}

	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("server shutting down")
	os.Exit(0)
}

func middleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestlog {
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
		w.Header().Set("Content-Type ", "text/xml; charset=utf-8")
		soapaction := r.Header.Get("Soapaction")
		switch soapaction {
		case "\"http://tempuri.org/ICreditCardService/CreditCardContracts\"":
			handlers.CreditCardContractsHandler(w, r)
		case "\"http://tempuri.org/ICreditCardService/ClientDataUpdate\"":
			handlers.ClientUpdateHandler(w, r)
		case "\"http://tempuri.org/ICreditCardService/CreditCardReject\"":
			handlers.CreditClaimRejectHandler(w, r)
		case "\"http://tempuri.org/ICreditCardService/CreditCardAccept\"":
			handlers.CreditClaimAcceptHandler(w, r)
		case "\"http://tempuri.org/ICreditCardService/CardClaimStatus\"":
			handlers.ClaimStatusHandler(w, r)
		case "\"http://tempuri.org/ICreditCardService/CreateCreditCardClaim\"":
			handlers.CreateCardClaimRHandler(w, r)
		}
	}
}
