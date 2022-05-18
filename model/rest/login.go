package rest

import (
	service "ec2/model/banco"
	"ec2/model/loggers"
	usuario "ec2/model/modelos"
	repositorio "ec2/model/repository"
	"ec2/model/validate"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Login(w, r)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user usuario.Usuario
	body, err := ioutil.ReadAll(r.Body)
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

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro conectar no db")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	repository := repositorio.NewRepositorio(db)
	userdb, err := repository.Login(user.Email)
	if err != nil {
		err = errors.New("erro email nao cadastrado")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	login := validate.CheckPasswordHash(user.Password, userdb.Password)
	if login != nil {
		err = errors.New("erro ao fazer login: senha incorreta")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	loggers.ResponseJson(w, http.StatusOK, "login feito com sucesso")
}