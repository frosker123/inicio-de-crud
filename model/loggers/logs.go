package loggers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Errors struct {
	Erros string `json:"Erros"`
}

func ResponseJson(w http.ResponseWriter, statuscode int, dados interface{}) {
	w.WriteHeader(statuscode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

func ResponseErrors(w http.ResponseWriter, statuscode int, err error) {
	ResponseJson(w, statuscode, Errors{
		Erros: err.Error(),
	})

}
