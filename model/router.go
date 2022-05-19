package main

import (
	congif "ec2/model/config"
	"ec2/model/midllewares"
	"ec2/model/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	congif.VariveisAm()

	//usuario
	router.HandleFunc("/usuario", rest.HandlerUsuario)
	router.HandleFunc("/usuario/put/{id}", rest.HandlerUsuario)
	router.HandleFunc("/usuario/{id}", midllewares.Authentic(rest.BuscaUsuarioById))

	//login
	router.HandleFunc("/login", rest.HandlerLogin)

	fmt.Printf("api rodando na porta %v\n", congif.API_PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", congif.API_PORT), router))
}
