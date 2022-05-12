package main

import (
	"ec2/model/rest"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuario", rest.HandlerUsuario)
	router.HandleFunc("/usuario/{id}", rest.HandlerUsuario)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("executando na porta %v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v", port), router))

}
