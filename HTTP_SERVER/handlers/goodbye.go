package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(m *log.Logger) *Goodbye {
	return &Goodbye{m}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Print("Hello Good\n")
	fmt.Fprint(rw, "Good byer")
}
