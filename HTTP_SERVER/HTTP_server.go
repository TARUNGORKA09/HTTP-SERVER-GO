package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/TARUNGORKA09/HTTP-SERVER-GO/HTTP_SERVER/handlers"
)

func main() {

	l := log.New(os.Stdout, "Product-API", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)
	prod := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/Goodbye", gb)
	sm.Handle("/product", prod)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()

	//tc, _ := context.WithDeadline(context.Background(), 10*time.Now()
	//s.Shutdown(t)
}
