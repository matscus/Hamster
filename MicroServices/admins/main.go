package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"context"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/admins/handlers"
	"github.com/matscus/Hamster/Package/Middleware/middleware"
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
	flag.StringVar(&listenport, "port", "10005", "port to Listen")
	flag.StringVar(&proto, "proto", "https", "http or https")
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
		Handler:      r,
	}

	r.HandleFunc("/api/v1/admins/getallusers", middleware.AdminsMiddleware(handlers.GetAllUsers)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/users", middleware.AdminsMiddleware(handlers.Users)).Methods(http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/changepassword", middleware.ChPassMiddleware(handlers.ChangePassword)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/getallhosts", middleware.AdminsMiddleware(handlers.GetAllHosts)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/getallhostswithproject", middleware.AdminsMiddleware(handlers.GetAllHostsWithProject)).Methods(http.MethodGet, http.MethodOptions).Queries("project", "{project}")
	r.HandleFunc("/api/v1/admins/hosts", middleware.AdminsMiddleware(handlers.Hosts)).Methods(http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/getallprojects", middleware.AdminsMiddleware(handlers.GetAllProjects)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/projects", middleware.AdminsMiddleware(handlers.Projects)).Methods(http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/roles", middleware.AdminsMiddleware(handlers.Roles)).Methods(http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/getallroles", middleware.AdminsMiddleware(handlers.GetAllRoles)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/servicebins", middleware.AdminsMiddleware(handlers.ServiceBins)).Methods(http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/getallservicebins", middleware.AdminsMiddleware(handlers.GetAllServiceBins)).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/admins/getallservicebinstype", middleware.AdminsMiddleware(handlers.GetAllServiceBinsType)).Methods(http.MethodGet, http.MethodOptions)
	r.PathPrefix("/api/v1/admins/bins/").Handler(http.StripPrefix("/api/v1/admins/files/", middleware.MiddlewareFiles(http.FileServer(http.Dir("home/matscus/Hamster/bins/"))))).Methods(http.MethodGet, http.MethodOptions).Queries("servicetype", "{servicetype}", "servicename", "{servicename}")
	http.Handle("/api/v1/admins/", r)
	//r.Use(mux.CORSMethodMiddleware(r))
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
