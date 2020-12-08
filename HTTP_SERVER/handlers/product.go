package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TARUNGORKA09/HTTP-SERVER-GO/HTTP_SERVER/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(m *log.Logger) *Products {
	return &Products{m}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshall ", http.StatusInternalServerError)
	}
	rw.Write(d)
}
