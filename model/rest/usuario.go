package rest

import (
	service "ec2/model/banco"
	usuario "ec2/model/modelos"
	"ec2/model/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

		w.Write([]byte("falha ao ler body da request"))
		return
	}
	err = json.Unmarshal(body, &usuario)
	if err != nil {

		w.Write([]byte("erro ao fazer unmarshal do usuario"))
		return
	}
	db, err := service.ConectaDB()
	if err != nil {

		w.Write([]byte("erro ao fazer conexão com o banco de dados"))
		return
	}

	repository := repository.NewRepositorio(db)
	_, err = repository.Criar(usuario)
	if err != nil {

		w.Write([]byte("erro ao criar usuario"))
		return
	}

	defer db.Close()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("novo usuario inserido com sucesso")))

}

func BuscaUsuarioById(w http.ResponseWriter, r *http.Request) {
	paramentros := mux.Vars(r)
	var usuario usuario.Usuario

	id, err := strconv.ParseUint(paramentros["id"], 10, 64)
	if err != nil {

		w.Write([]byte("erro ao achar id"))
		return
	}
	db, err := service.ConectaDB()
	if err != nil {

		w.Write([]byte("erro ao fazer conexão com o banco de dados"))
		return
	}

	row, err := db.Query("select * from usuarios.usuarios where id = $1 ", id)
	if err != nil {

		w.Write([]byte("erro, id nao encontrado ou nao existe"))
		return
	}

	if row.Next() {
		if err := row.Scan(&usuario.ID, &usuario.Nome, &usuario.UserName, &usuario.Email, &usuario.Password); err != nil {

			w.Write([]byte("erro ao retornar usuario o scan pra mostrar o usuario"))
			return
		}
	}

	if usuario.ID == 0 {

		w.Write([]byte("erro, id nao encontrado ou nao existe"))
		return
	}

	defer db.Close()

	err = json.NewEncoder(w).Encode(usuario)
	if err != nil {

		w.Write([]byte("erro ao converter usuario para json"))
		return
	}

}
