package repositorio

import (
	"database/sql"
	usuario "ec2/model/modelos"
	"fmt"
)

type repositorio struct {
	db *sql.DB
}

func NewRepositorio(db *sql.DB) *repositorio {
	return &repositorio{db}
}

func (repo repositorio) Criar(user usuario.Usuario) (int64, error) {
	statement := `insert into usuarios.usuarios(nome, username, email, password, created_at)values($1, $2, $3, $4,$5)`
	_, err := repo.db.Exec(statement, user.Nome, user.UserName, user.Email, user.Password, user.DataCriacao)
	if err != nil {
		panic(err)
	}
	return user.ID, nil
}

func (repo repositorio) Querie(nikeouName string) ([]usuario.Usuario, error) {
	nikeouName = fmt.Sprintf("%%s%%")
	var usuarios []usuario.Usuario

	row, err := repo.db.Query("select id, nome, username, email, created_at from usuarios.usuarios where nome like $1 or username like $2", nikeouName, nikeouName)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var usuario usuario.Usuario
		if err = row.Scan(&usuario.ID, &usuario.Nome, &usuario.UserName, &usuario.Email, &usuario.DataCriacao); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repo repositorio) GetbyId(id int64) (usuario.Usuario, error) {
	var user usuario.Usuario

	row, err := repo.db.Query("select id, nome, username, email, created_at from usuarios.usuarios where id = $1", id)
	if err != nil {
		return usuario.Usuario{}, err
	}

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Nome,
			&user.UserName,
			&user.Email,
			&user.DataCriacao,
		); err != nil {
			return usuario.Usuario{}, nil
		}
	}

	if user.ID == 0 {
		return usuario.Usuario{}, err

	}

	return user, nil
}
