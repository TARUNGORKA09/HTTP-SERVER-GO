package handlers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/TARUNGORKA09/HTTP-SERVER-GO/HTTP_SERVER/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(m *log.Logger) *Products {
	return &Products{m}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	}
	//handle the update
	if r.Method == http.MethodPost {
		p.addproduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		//except the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall ", http.StatusInternalServerError)
	}

}
func (p *Products) addproduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Print("Update POST method")

	prod := &data.Product{}
	err := prod.ToDATA(r.Body)
	if err != nil {
		http.Error(rw, "unable to Convert ", http.StatusInternalServerError)
	}
	p.l.Print("Data %#v", prod)
	data.AddProduct(prod)
}
