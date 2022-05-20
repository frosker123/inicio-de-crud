package router

import (
	"github.com/gorilla/mux"
)

func GerarAPI() *mux.Router {
	r := mux.NewRouter()
	routers := ConfiguraRotas(r)
	return routers
}
