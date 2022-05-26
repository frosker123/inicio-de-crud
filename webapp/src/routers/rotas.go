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
	rotas := configRotas

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Handlerfunc).Methods(rota.Metodo)
	}

	//bloco que carregas os arquivos de estilo do css
	listeServ := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", listeServ))

	return r
}
