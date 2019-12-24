package main

import (
	"context"
	"flag"
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
)

func main() {
	flag.StringVar(&pemPath, "pempath", os.Getenv("SERVERREM"), "path to pem file")
	flag.StringVar(&keyPath, "keypath", os.Getenv("SERVERKEY"), "path to key file")
	flag.StringVar(&listenport, "port", "10000", "port to Listen")
	flag.StringVar(&proto, "proto", "http", "http or https")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully")
	flag.DurationVar(&readTimeout, "read-timeout", time.Second*15, "read server timeout")
	flag.DurationVar(&writeTimeout, "write-timeout", time.Second*15, "write server timeout")
	flag.DurationVar(&idleTimeout, "idle-timeout", time.Second*60, "idle server timeout")
	flag.Parse()
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         "127.0.0.1:" + listenport,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
	r.HandleFunc("/ICreditCardService/CardClaimStatus", handlers.WcfCreditCardService).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
	r.HandleFunc("/ICreditCardService/ClientDataUpdate", handlers.WcfCreditCardService).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
	r.HandleFunc("/ICreditCardService/CreateCreditCardClaim", handlers.WcfCreditCardService).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
	r.HandleFunc("/ICreditCardService/CreditCardAccept", handlers.WcfCreditCardService).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
	r.HandleFunc("/ICreditCardService/CreditCardContracts", handlers.WcfCreditCardService).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
	r.HandleFunc("/ICreditCardService/CreditCardReject", handlers.WcfCreditCardService).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
	r.HandleFunc("/WcfCreditCardService/CreditCardService.svc", handlers.WcfCreditCardService).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
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
