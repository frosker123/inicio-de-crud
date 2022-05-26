package router

import "github.com/gorilla/mux"

func GerarAPI() *mux.Router {
	r := mux.NewRouter()
	router := ConfiguraRotas(r)

	return router
}
