package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibnuda/sempak/src/handlers"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "static", "default thingy.")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/sempak/{ukuran:[0-9]+}", handlers.EmbuhHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
