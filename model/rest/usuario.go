package rest

import (
	service "ec2/model/banco"
	"ec2/model/loggers"
	usuario "ec2/model/modelos"
	repositorio "ec2/model/repository"
	"ec2/model/validate"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func HandlerUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		InserirUsuario(w, r)
	}

	if r.Method == http.MethodGet {
		BuscaUsuario(w, r)
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

	repository := repositorio.NewRepositorio(db)
	_, err = repository.Criar(usuario)
	if err != nil {
		err = errors.New("erro ao criar usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	defer db.Close()

	loggers.ResponseJson(w, http.StatusOK, usuario)

}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	nikeouName := strings.ToLower(r.URL.Query().Get("user"))

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro conectar no db")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	repository := repositorio.NewRepositorio(db)
	querier, err := repository.Querie(nikeouName)
	if err != nil {
		err = errors.New("erro no filtro de usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	loggers.ResponseJson(w, http.StatusOK, querier)

}

func BuscaUsuarioById(w http.ResponseWriter, r *http.Request) {
	paramentros := mux.Vars(r)

	usuarioid := paramentros["id"]

	id, err := strconv.ParseInt(usuarioid, 0, 64)
	if err != nil {
		err = errors.New("erro ao converter id ")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro ao conectar no banco")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}
	repository := repositorio.NewRepositorio(db)
	buscabyID, err := repository.GetbyId(id)
	if err != nil {
		err = errors.New("erro ao achar id do usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	defer db.Close()

	loggers.ResponseJson(w, http.StatusOK, buscabyID)

}
