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

	if r.Method == http.MethodPut {
		PutUsuario(w, r)
	}

}

func InserirUsuario(w http.ResponseWriter, r *http.Request) {
	var user usuario.Usuario

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = errors.New("erro ao ler o body")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		err = errors.New("erro ao fazer unmarshal")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	if err = validate.Valid(&user, "inserir"); err != nil {
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
	_, err = repository.Create(user)
	if err != nil {
		err = errors.New("erro ao criar usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	defer db.Close()

	loggers.ResponseJson(w, http.StatusOK, user)

}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	nikeouName := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro conectar no db")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	repository := repositorio.NewRepositorio(db)
	querier, err := repository.GetUser(nikeouName)
	if err != nil {
		err = errors.New("erro no filtro de usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	if querier == nil {
		err = errors.New("usuario/s informado nao encontrado")
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

	if buscabyID.ID == 0 {
		err = errors.New("id informado nao existe")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	defer db.Close()

	loggers.ResponseJson(w, http.StatusOK, buscabyID)

}

func PutUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	att := usuario.Usuario{
		Nome:     r.FormValue("nome"),
		UserName: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	id := params["id"]

	idUser, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		err = errors.New("erro ao converter id para um int")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro ao conectar no banco")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	if err = validate.Valid(&att, "atualiza"); err != nil {
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	repository := repositorio.NewRepositorio(db)
	err = repository.AttUser(idUser, att)
	if err != nil {
		err = errors.New("erro ao atualizar usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	loggers.ResponseJson(w, http.StatusOK, "usuario atualizado :)")
}
