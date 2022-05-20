package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
	URI          string
	Metodo       string
	Handlerfunc  func(w http.ResponseWriter, r *http.Request)
	Autenticacao bool
}

func ConfiguraRotas(r *mux.Router) *mux.Router {
	rotas := rotaConfig

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Handlerfunc).Methods(rota.Metodo)
	}

	return r
}
