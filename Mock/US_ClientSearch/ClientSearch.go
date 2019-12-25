package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	file []byte
	err  error
)

func handler(w http.ResponseWriter, r *http.Request) {
	traceHeader := r.Header.Get("GPB-traceId")
	if traceHeader != "" {
		w.Header().Set("GPB-traceId", traceHeader)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(file)
}

func middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("[INFO] %s %s", r.URL, time.Since(start))
		}()
		f(w, r)
	}
}

func init() {
	file, err = ioutil.ReadFile("./response.json")
	if err != nil {
		log.Println("[ERROR] can't open file ./response.json")
	}
}

func main() {
	mode := flag.String("mode", "http", "http or https")
	port := flag.String("port", "9999", "listen port")
	endpoint := flag.String("endpoint", "/", "endpoint")
	pem := flag.String("pem", "./ssl/server.pem", "pem path")
	key := flag.String("key", "./ssl/server.key", "key path")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc(*endpoint, handler)
	srv := &http.Server{
		Addr:         "0.0.0.0:" + *port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	log.Printf("[INFO] Starting mock %s://localhost:%s%s", *mode, *port, *endpoint)
	switch *mode {
	case "http":
		log.Fatal(srv.ListenAndServe())
	case "https":
		log.Fatal(srv.ListenAndServeTLS(*pem, *key))
	}
}
