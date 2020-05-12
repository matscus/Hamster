package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/matscus/Hamster/Mock/info_service/datapool"
	"github.com/matscus/Hamster/Mock/info_service/handlers"

	"github.com/gorilla/mux"
)

var (
	pemPath      string
	keyPath      string
	proto        string
	listenport   string
	host         string
	wait         time.Duration
	writeTimeout time.Duration
	readTimeout  time.Duration
	idleTimeout  time.Duration
)

func init() {
	go datapool.IntitDataPool()
}

func main() {
	flag.StringVar(&pemPath, "pempath", os.Getenv("SERVERREM"), "path to pem file")
	flag.StringVar(&keyPath, "keypath", os.Getenv("SERVERKEY"), "path to key file")
	flag.StringVar(&listenport, "port", "9092", "port to Listen")
	flag.StringVar(&datapool.FilePath, "filepath", "datapool.csv", "path from csv data file")
	flag.StringVar(&proto, "proto", "http", "http or https")
	flag.BoolVar(&handlers.Requestlog, "request-log", false, "idle server timeout")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully")
	flag.DurationVar(&readTimeout, "read-timeout", time.Second*15, "read server timeout")
	flag.DurationVar(&writeTimeout, "write-timeout", time.Second*15, "write server timeout")
	flag.DurationVar(&idleTimeout, "idle-timeout", time.Second*60, "idle server timeout")
	flag.Float64Var(&handlers.Mean, "mean-timeout", 0.0, "mean responce timeout")
	flag.Float64Var(&handlers.Deviation, "deviation-timeout", 0, "deviation responce timeout")
	flag.Parse()

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				host = ipnet.IP.String()
			}
		}
	}

	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         host + ":" + listenport,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
	r.HandleFunc("/omni-information/api/v2/client/search", handlers.Middleware(handlers.ClientSearchCommon)).Methods(http.MethodPost)
	r.HandleFunc("/omni-information/api/v2/client/product/deposit/list", handlers.Middleware(handlers.DepositList)).Methods(http.MethodPost)
	r.HandleFunc("/omni-information/api/v2/client/product/account/list", handlers.Middleware(handlers.AccountList)).Methods(http.MethodPost)
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
