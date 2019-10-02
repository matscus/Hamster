package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/Mock/dadata/cache"
	"github.com/matscus/Hamster/Mock/dadata/handlers"
)

func init() {
	cache.LoadCache()
}

var (
	listenport string
	mode       string
)

func main() {
	flag.StringVar(&mode, "mode", "https", "server mode")
	flag.StringVar(&listenport, "port", ":9443", "port to Listen")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/dadata/search/fio", handlers.GetFIO).Methods("POST")
	r.HandleFunc("/api/v1/dadata/search/fio", handlers.GetFIO).Methods("GET").Queries("meta_chanel", "{meta_chanel}", "query", "{query}")
	r.HandleFunc("/api/v1/dadata/search/address", handlers.GetAddress).Methods("POST")
	r.HandleFunc("/api/v1/dadata/search/address", handlers.GetAddress).Methods("GET").Queries("meta_chanel", "{meta_chanel}", "query", "{query}")
	r.HandleFunc("/api/v1/dadata/search/party", handlers.GetOrganization).Methods("POST")
	r.HandleFunc("/api/v1/dadata/search/party", handlers.GetOrganization).Methods("GET").Queries("meta_chanel", "{meta_chanel}", "query", "{query}")
	http.Handle("/api/v1/dadata/search/", r)
	switch mode {
	case "http":
		log.Printf("Listen to http on port %s", listenport)
		log.Fatal(http.ListenAndServe(listenport, r))
	case "https":
		log.Printf("Listen to https on port %s", listenport)
		log.Fatal(http.ListenAndServeTLS(listenport, os.Getenv("SERVERREM"), os.Getenv("SERVERKEY"), context.ClearHandler(http.DefaultServeMux)))
	}
}
