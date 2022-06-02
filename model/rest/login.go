package rest

import (
	service "ec2/model/banco"
	"ec2/model/loggers"
	usuario "ec2/model/modelos"
	repositorio "ec2/model/repository"
	"ec2/model/token"
	"ec2/model/validate"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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

	err = validate.CheckPasswordHash(userdb.Password, user.Password)
	if err != nil {
		err = errors.New("erro senhas nao batem")
		loggers.ResponseErrors(w, http.StatusInternalServerError, err)
		return
	}

	token, err := token.CreateToken(userdb.ID)
	if err != nil {
		err = errors.New("erro ao gerar token para usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	loggers.ResponseText(w, http.StatusOK, token)

}
