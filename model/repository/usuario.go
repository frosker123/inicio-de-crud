package repository

import (
	"database/sql"
	usuario "ec2/model/modelos"
)

type Usuario struct {
	db *sql.DB
}

func NewRepositorio(db *sql.DB) *Usuario {
	return &Usuario{db}
}

func (repository Usuario) Criar(user usuario.Usuario) (int64, error) {
	statement := `insert into usuarios.usuarios(nome, username, email, password)values($1, $2, $3, $4)`
	_, err := repository.db.Exec(statement, user.Nome, user.UserName, user.Email, user.Password)
	if err != nil {
		panic(err)
	}
	return user.ID, nil
}
