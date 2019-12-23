package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/matscus/Hamster/MicroServices/scenario/handlers"
	"github.com/matscus/Hamster/Package/Middleware/middleware"

	"context"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
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

func init() {
	go func() {
		for {
			scn.InitData()
			time.Sleep(1 * time.Minute)
		}
	}()
}

func main() {
	flag.StringVar(&pemPath, "pempath", os.Getenv("SERVERREM"), "path to pem file")
	flag.StringVar(&keyPath, "keypath", os.Getenv("SERVERKEY"), "path to key file")
	flag.StringVar(&listenport, "port", "10004", "port to Listen")
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
	r.HandleFunc("/api/v1/scenario/start", middleware.Middleware(handlers.StartScenario)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/scenario/stop", middleware.Middleware(handlers.StopScenario)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/scenario/new", middleware.Middleware(handlers.NewScenario)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/v1/scenario", middleware.Middleware(handlers.GetScenarios)).Methods(http.MethodGet, http.MethodOptions).Queries("project", "{project}")
	r.HandleFunc("/api/v1/scenario", middleware.Middleware(handlers.UpdateOrDeleteScenario)).Methods(http.MethodPut, http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/api/v1/scenario/lastparams", middleware.Middleware(handlers.GetLastParams)).Methods(http.MethodGet, http.MethodOptions).Queries("name", "{name}", "project", "{project}")
	r.HandleFunc("/api/v1/scenario/ws", handlers.Ws)
	r.PathPrefix("/api/v1/scenario/files/").Handler(http.StripPrefix("/api/v1/scenario/files/", handlers.MiddlewareFiles(http.FileServer(http.Dir("/home/matscus/Hamster/projects/"))))).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/v1/scenario/precheck", middleware.Middleware(handlers.PreCheckScenario)).Methods(http.MethodPost, http.MethodOptions)
	http.Handle("/api/v1/", r)
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
