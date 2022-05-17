package rest

import (
	service "ec2/model/banco"
	"ec2/model/loggers"
	usuario "ec2/model/modelos"
	"ec2/model/repository"
	"ec2/model/validate"
	"encoding/json"
	"errors"
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
		err = errors.New("erro ao ler o body")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(body, &usuario)
	if err != nil {
		err = errors.New("erro ao fazer unmarshal")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	if err = validate.Valid(&usuario); err != nil {
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro conectar no db")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	repository := repository.NewRepositorio(db)
	_, err = repository.Criar(usuario)
	if err != nil {
		err = errors.New("erro ao criar usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	defer db.Close()

	loggers.ResponseJson(w, http.StatusOK, fmt.Sprintf("usuario inserido com sucesso %v", usuario))

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
