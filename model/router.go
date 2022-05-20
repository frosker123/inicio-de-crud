package main

import (
	congif "ec2/model/config"
	"ec2/model/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	router.GerarAPI()
	congif.VariveisAm()

	fmt.Printf("api rodando na porta %v\n", congif.API_PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", congif.API_PORT), router.GerarAPI()))
}
