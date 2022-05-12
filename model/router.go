package main

import (
	"ec2/model/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuario", rest.HandlerUsuario)

	port := "8000"
	if port == "" {
		port = "8080"
	}

	fmt.Printf("executando na porta %v", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))

}
