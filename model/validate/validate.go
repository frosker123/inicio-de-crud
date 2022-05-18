package validate

import (
	usuario "ec2/model/modelos"
	"errors"
	"regexp"
	"strings"
	"time"
)

func Valid(user *usuario.Usuario) error {
	if err := validateCampo(user); err != nil {
		return err
	}
	formate(user)

	return nil
}

func validateCampo(user *usuario.Usuario) error {
	if user.Nome == "" {
		return errors.New("campo nome tem que ser preenchido")
	}

	if user.UserName == "" {
		return errors.New("campo user name tem que ser preenchido")
	}

	if user.Email == "" {
		return errors.New("campo email tem que ser preenchido")
	}

	err := IsEmailValid(user.Email)
	if !err {
		return errors.New("o email informado é invalido")
	}

	if user.Password == "" {
		return errors.New("campo senha tem que ser preenchido")
	}

	user.DataCriacao = time.Now()

	return nil
}

func formate(user *usuario.Usuario) {
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)
	user.Nome = strings.TrimSpace(user.Nome)

}

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
