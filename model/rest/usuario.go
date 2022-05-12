package rest

import (
	service "ec2/model/services"
	"ec2/model/usuario"
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

	defer db.Close()

	sqlStatement, err := db.Prepare(`INSERT INTO usuarios.usuario (nome, email)VALUES ($1, $2)`)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("erro ao fazer o statement"))
		return
	}

	defer sqlStatement.Close()

	insertline, err := sqlStatement.Exec(sqlStatement, usuario.Nome, usuario.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("erro ao fazer o inserir dado"))
		return
	}

	inserirId, err := insertline.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("erro ao fazer o inserir id"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("novo usuario inserido com id %v", inserirId)))

}

func BuscaUsuarioById(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	var usuario usuario.Usuario

	id, err := strconv.ParseInt(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("erro ao mostrar id usuario"))
		return
	}

	db, err := service.ConectaDB()
	if err != nil {
		w.Write([]byte("erro ao fazer conexão com o banco de dados"))
		return
	}

	query, err := db.Query("select * from usuario id = ?", id)
	if err != nil {
		w.Write([]byte("erro ao buscar user"))
	}

	if query.Next() {
		if err := query.Scan(&usuario.ID, &usuario.Nome, usuario.Email); err != nil {
			w.Write([]byte("erro ao fazer scan do usuario"))
		}
	}

	if err := json.NewEncoder(w).Encode(&usuario); err != nil {
		w.Write([]byte("erro ao fazer encode"))
	}
}
