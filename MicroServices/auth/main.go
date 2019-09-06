package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/matscus/Hamster/MicroServices/auth/handlers"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	flag.StringVar(&listenport, "port", ":10000", "port to Listen")
	flag.StringVar(&proto, "proto", "https", "http or https")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/auth/new", handlers.Middleware(handlers.GetToken)).Methods("POST", "OPTIONS")
	http.Handle("/api/v1/auth/", r)
	log.Println("ListenAndServe: " + listenport)
	err := http.ListenAndServeTLS(listenport, pemPath, keyPath, context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
