package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/matscus/Hamster/MicroServices/service/handlers"
	"github.com/matscus/Hamster/MicroServices/service/serv"
	"github.com/matscus/Hamster/Package/Middleware/middleware"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var (
	pemPath    string
	keyPath    string
	proto      string
	listenport string
)

//Token - struct for auth token
type Token struct {
	Token string `json:"token"`
}

func init() {
	go serv.InitGetResponseAllData()
	serv.CheckService()
}

func main() {
	flag.StringVar(&pemPath, "pem", os.Getenv("SERVERREM"), "path to pem file")
	flag.StringVar(&keyPath, "key", os.Getenv("SERVERKEY"), "path to key file")
	flag.StringVar(&listenport, "port", ":10001", "port to Listen")
	flag.StringVar(&proto, "proto", "https", "http or https")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/service/start", middleware.Middleware(handlers.StartSeviceHandle)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service/stop", middleware.Middleware(handlers.StopServiceHandle)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service", middleware.Middleware(handlers.GetServicesHandle)).Methods("GET", "OPTIONS").Queries("project", "{project}")
	r.HandleFunc("/api/v1/service/update", middleware.Middleware(handlers.UpdateServiceHandle)).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/service/new", middleware.Middleware(handlers.NewServiceHandle)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service/delete", middleware.Middleware(handlers.DeleteService)).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/v1/service/generators", middleware.Middleware(handlers.NewOrUpdateGenerator)).Methods("GET", "POST", "PUT", "OPTIONS")
	r.HandleFunc("/api/v1/service/newfile", middleware.Middleware(handlers.NewFile)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service/file", middleware.Middleware(handlers.File)).Methods("GET", "DELETE", "OPTIONS").Queries("servicetype", "{servicetype}", "servicename", "{servicename}")
	http.Handle("/api/v1/", r)
	log.Println("ListenAndServe: " + listenport)
	err := http.ListenAndServeTLS(listenport, pemPath, keyPath, context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
