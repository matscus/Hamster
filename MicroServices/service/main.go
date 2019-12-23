package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/matscus/Hamster/MicroServices/service/handlers"
	"github.com/matscus/Hamster/MicroServices/service/serv"
	"github.com/matscus/Hamster/Package/Middleware/middleware"

	"context"

	"github.com/gorilla/mux"
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
	r.HandleFunc("/api/v1/service/start", middleware.Middleware(handlers.StartSevice)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service/stop", middleware.Middleware(handlers.StopService)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service", middleware.Middleware(handlers.GetAllServices)).Methods("GET", "OPTIONS").Queries("project", "{project}")
	r.HandleFunc("/api/v1/service/update", middleware.Middleware(handlers.UpdateServiceHandle)).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/v1/service/new", middleware.Middleware(handlers.NewServiceHandle)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service/delete", middleware.Middleware(handlers.DeleteService)).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/v1/service/newfile", middleware.Middleware(handlers.NewFile)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/v1/service/file", middleware.Middleware(handlers.File)).Methods("GET", "DELETE", "OPTIONS").Queries("servicetype", "{servicetype}", "servicename", "{servicename}")
	http.Handle("/api/v1/", r)
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
