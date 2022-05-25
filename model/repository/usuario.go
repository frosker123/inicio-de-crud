package repositorio

import (
	"database/sql"
	usuario "ec2/model/modelos"
	"errors"
	"fmt"
)

type repositorio struct {
	db *sql.DB
}

func NewRepositorio(db *sql.DB) *repositorio {
	return &repositorio{db}
}

func (repo repositorio) Create(user usuario.Usuario) (int64, error) {
	statement := `insert into usuarios.usuarios(nome, username, email, password, created_at)values($1, $2, $3, $4,$5)`
	_, err := repo.db.Exec(statement, user.Nome, user.UserName, user.Email, user.Password, user.DataCriacao)
	if err != nil {
		return 0, errors.New("erro na update de usuario ")
	}
	return user.ID, nil
}

func (repo repositorio) GetUser(nikeouName string) ([]usuario.Usuario, error) {
	nikeouName = fmt.Sprintf("%%%s%%", nikeouName)
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

func (repo repositorio) GetPass(id int64) (string, error) {
	var user usuario.Usuario

	row, err := repo.db.Query("select password from usuarios.usuarios where id = $1", id)
	if err != nil {
		return "", errors.New("erro na query de senha ")
	}

	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", errors.New("erro no scan de senha ")
		}
	}

	return user.Password, nil
}

func (repo repositorio) AttPass(id int64, senha string) error {
	statement := `update usuarios.usuarios set password = $1 where id = $2`
	_, err := repo.db.Exec(statement, senha, id)
	if err != nil {
		return errors.New("erro na update")
	}

	return nil
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

	return user, nil
}

func (repo repositorio) AttUser(id int64, user usuario.Usuario) error {

	statement := `update usuarios.usuarios set nome = $1, username = $2, email = $3 where id = $4`
	_, err := repo.db.Exec(statement, user.Nome, user.UserName, user.Email, id)
	if err != nil {
		return errors.New("erro na update")
	}

	return nil
}

func (repo repositorio) Login(email string) (usuario.Usuario, error) {
	var user usuario.Usuario

	row, err := repo.db.Query("select id, password from usuarios.usuarios where email = $1", email)
	if err != nil {
		return usuario.Usuario{}, err
	}

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return usuario.Usuario{}, nil
		}
	}

	return user, nil
}
