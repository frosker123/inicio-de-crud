package main

import (
	"fmt"
	"log"
	"net/http"
	router "webapp/src/routers"
	"webapp/utils"
)

func main() {
	router := router.GerarAPI()
	utils.CarregaTemplate()

	fmt.Println("rodando api ")
	log.Fatal(http.ListenAndServe(":3000", router))
}
