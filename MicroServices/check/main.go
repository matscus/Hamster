package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/matscus/Hamster/MicroServices/check/handlers"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var (
	pemPath    string
	keyPath    string
	proto      string
	listenport string
)

func main() {
	flag.StringVar(&pemPath, "pem", os.Getenv("SERVERREM"), "path to pem file")
	flag.StringVar(&keyPath, "key", os.Getenv("SERVERKEY"), "path to key file")
	flag.StringVar(&listenport, "port", ":10003", "port to Listen")
	flag.StringVar(&proto, "proto", "https", "http or https")
	flag.Parse()
	log.SetFlags(1)
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/check/ws", handlers.Ws)
	http.Handle("/api/v1/check/", r)
	log.Println("ListenAndServe: " + listenport)
	err := http.ListenAndServeTLS(listenport, pemPath, keyPath, context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
