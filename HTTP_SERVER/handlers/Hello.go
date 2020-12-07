package handlers

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
} 

func (h *Hello) ServeHTTP(rw http.ResponseWriter,r *http.Request){
	h.l.Print("Hello John\n")
	d, err := ioutil.ReadAll(r.Body)
	if err!= nil{
		http.Error(w,"Oops", http.StatusBadRequest)
		return
	}
	log.Print("Data %s\n", d)
	fmt.Fprint(rw, "HELLO %s", d)
}