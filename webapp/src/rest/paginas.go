package rest

import (
	"net/http"
	"webapp/utils"
)

func TelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

func CarregarTelaCriarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "cadastro.html", nil)
}
