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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

func ResponseErrors(w http.ResponseWriter, statuscode int, err error) {
	ResponseJson(w, statuscode, Errors{
		Erros: err.Error(),
	})

}

func ResponseText(w http.ResponseWriter, statuscode int, text string) {
	if text != "" {
		w.Write([]byte(text))
	}
}
