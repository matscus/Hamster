package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/Mock/contract_service/handlers"
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
	Mean         float64
	Deviation    float64
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
	flag.Float64Var(&Mean, "mean-timeout", 0.0, "mean responce timeout")
	flag.Float64Var(&Deviation, "deviation-timeout", 0, "deviation responce timeout")
	flag.Parse()
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         "0.0.0.0:" + listenport,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
	r.HandleFunc("/WcfLoanIssueService/LoanIssue.svc", middleware()).Methods(http.MethodPost, http.MethodOptions, http.MethodGet)
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
		if Mean != 0.0 {
			waitResponse()
		}
		handlers.IssueComplete(w, r)
	}
}

func waitResponse() {
	timeout := rand.NormFloat64()*Deviation + Mean
	time.Sleep(time.Duration(timeout) * time.Millisecond)
}
