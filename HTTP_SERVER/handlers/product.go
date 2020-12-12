package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TARUNGORKA09/HTTP-SERVER-GO/HTTP_SERVER/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(m *log.Logger) *Products {
	return &Products{m}
}

type keyProduct struct {
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	lp := data.GetProduct(id)
	err1 := lp.ToJSON(rw)
	if err1 != nil {
		http.Error(rw, "unable to marshall ", http.StatusInternalServerError)
	}

}
func (p *Products) Addproducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Print("Update POST method")

	prod := r.Context().Value(keyProduct{}).(data.Product)
	p.l.Print("Data %#v", prod)
	data.AddProduct(&prod)
}
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(keyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
}

func (P Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err1 := prod.ToDATA(r.Body)
		if err1 != nil {
			http.Error(rw, "unable to Unmarshall JSON ", http.StatusInternalServerError)
		}

		//validate the product
		err := prod.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprintf("Validating error %s", err), http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)

	})
}
