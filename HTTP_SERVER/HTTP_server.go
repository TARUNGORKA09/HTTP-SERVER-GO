package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	l := log.New(os.Stdout, "Product-API", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm = http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()
}
