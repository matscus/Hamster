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
	"github.com/matscus/Hamster/Mock/omni_glossary/handlers"
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
	flag.BoolVar(&handlers.Requestlog, "request-log", false, "idle server timeout")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully")
	flag.DurationVar(&readTimeout, "read-timeout", time.Second*15, "read server timeout")
	flag.DurationVar(&writeTimeout, "write-timeout", time.Second*15, "write server timeout")
	flag.DurationVar(&idleTimeout, "idle-timeout", time.Second*60, "idle server timeout")
	flag.Float64Var(&handlers.Mean, "mean-timeout", 0.0, "mean responce timeout")
	flag.Float64Var(&handlers.Deviation, "deviation-timeout", 0, "deviation responce timeout")
	flag.Parse()
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         "0.0.0.0:" + listenport,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
	r.HandleFunc("/omni-glossary/api/v1/office/list", handlers.Middleware(handlers.Search)).Methods(http.MethodPost)
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
