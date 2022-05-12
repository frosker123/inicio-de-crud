package rest

import (
	service "ec2/model/services"
	"ec2/model/usuario"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandlerUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		InserirUsuario(w, r)
	}

	if r.Method == http.MethodGet {
		BuscaUsuarioById(w, r)
	}

}

func InserirUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario usuario.Usuario

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("falha ao ler body da request"))
		return
	}

	err = json.Unmarshal(body, &usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("erro ao fazer unmarshal do usuario"))
		return
	}

	db, err := service.ConectaDB()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("erro ao fazer conexão com o banco de dados"))
		return
	}

	statement := `insert into usuarios.usuario(nome, email)values($1, $2)`
	_, e := db.Exec(statement, usuario.Nome, usuario.Email)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("erro ao inserir usuario"))
		return
	}

	defer db.Close()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("novo usuario inserido com sucesso")))

}

func BuscaUsuarioById(w http.ResponseWriter, r *http.Request) {

}
