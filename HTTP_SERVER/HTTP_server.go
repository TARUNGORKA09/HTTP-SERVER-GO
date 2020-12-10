package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TARUNGORKA09/HTTP-SERVER-GO/HTTP_SERVER/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "Product-API", log.LstdFlags)
	prod := handlers.NewProduct(l)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", prod.GetProducts)
	getRouter.Use(prod.MiddlewareValidateProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", prod.UpdateProducts)
	putRouter.Use(prod.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/addproduct", prod.Addproducts)
	postRouter.Use(prod.MiddlewareValidateProduct)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()

	//tc, _ := context.WithDeadline(context.Background(), 30*time.Second)
	//s.Shutdown(tc)
}
