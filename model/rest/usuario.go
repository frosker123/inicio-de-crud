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

	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

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
		loggers.ResponseErrors(w, http.StatusInternalServerError, err)
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
		loggers.ResponseErrors(w, http.StatusInternalServerError, err)
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
		loggers.ResponseErrors(w, http.StatusInternalServerError, err)
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

	id := params["id"]

	idUser, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		err = errors.New("erro ao converter id para um int")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	idToken, err := token.TokenIdUser(r)
	if err != nil {
		err = errors.New("erro no id do token")
		loggers.ResponseErrors(w, http.StatusUnauthorized, err)
		return
	}

	if idToken != idUser {
		err = errors.New("usuario nao pode atualizar esse perfil")
		loggers.ResponseErrors(w, http.StatusForbidden, err)
		return
	}

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro ao conectar no banco")
		loggers.ResponseErrors(w, http.StatusInternalServerError, err)
		return
	}

	if err = validate.Valid(&user, "atualiza"); err != nil {
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	repository := repositorio.NewRepositorio(db)
	err = repository.AttUser(idUser, user)
	if err != nil {
		err = errors.New("erro ao atualizar usuario")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	loggers.ResponseJson(w, http.StatusOK, "usuario atualizado :)")
}

func PutPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user usuario.Usuario

	idToken, err := token.TokenIdUser(r)
	if err != nil {
		err = errors.New("erro no id do token")
		loggers.ResponseErrors(w, http.StatusUnauthorized, err)
		return
	}

	id := params["id"]

	idUser, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		err = errors.New("erro ao converter id para um int")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	if idToken != idUser {
		err = errors.New("usuario nao pode atualizar esse perfil")
		loggers.ResponseErrors(w, http.StatusForbidden, err)
		return
	}

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

	db, err := service.ConectaDB()
	if err != nil {
		err = errors.New("erro ao conectar no banco")
		loggers.ResponseErrors(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorio.NewRepositorio(db)
	senhadoDB, err := repository.GetPass(idUser)
	if err != nil {
		err = errors.New("erro ao obter senha")
		loggers.ResponseErrors(w, http.StatusInternalServerError, err)
		return
	}

	err = validate.CheckPasswordHash(senhadoDB, user.Password)
	if err != nil {
		err = errors.New("senha nao é a mesma do banco")
		loggers.ResponseErrors(w, http.StatusUnauthorized, err)
		return
	}

	senhacomHash, err := validate.HashPassword(user.NewPassword)
	if err != nil {
		err = errors.New("erro ao criptografar senha ")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	err = repository.AttPass(idUser, string(senhacomHash))
	if err != nil {
		err = errors.New("erro ao atualizar senha")
		loggers.ResponseErrors(w, http.StatusBadRequest, err)
		return
	}

	loggers.ResponseText(w, http.StatusOK, "senha atualizada")

}
