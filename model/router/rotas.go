package router

import (
	"ec2/model/midllewares"
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
		if rota.Autenticacao {
			r.HandleFunc(rota.URI, midllewares.Authentic(rota.Handlerfunc)).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, rota.Handlerfunc).Methods(rota.Metodo)

		}
	}

	return r
}
