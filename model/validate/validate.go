package validate

import (
	usuario "ec2/model/modelos"
	"errors"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Valid(user *usuario.Usuario, valida string) error {
	if err := validateCampo(user, valida); err != nil {
		return err
	}
	err := formate(valida)
	if err != nil {
		return err
	}

	return nil
}

func validateCampo(user *usuario.Usuario, valida string) error {

	if user.Nome == "" {
		return errors.New("campo nome tem que ser preenchido")
	}

	if user.UserName == "" {
		return errors.New("campo user name tem que ser preenchido")
	}

	if user.Email == "" {
		return errors.New("campo email tem que ser preenchido")
	}

	if valida == "inserir" {
		err := IsEmailValid(user.Email)
		if !err {
			return errors.New("o email informado é invalido")
		}
	}

	if valida == "inserir" && user.Password == "" {
		return errors.New("campo senha tem que ser preenchido")
	}

	senha, e := HashPassword(user.Password)
	if e != nil {
		return errors.New("senha nao  criptografada ")
	}

	user.Password = string(senha)

	user.DataCriacao = time.Now()

	return nil
}

func formate(valida string) error {
	var user usuario.Usuario
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)
	user.Nome = strings.TrimSpace(user.Nome)

	if valida == "inserir" {
		newSenha, err := HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = newSenha
	}

	return nil
}

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func HashPassword(password string) (string, error) {
	senha, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(senha), nil
}

func CheckPasswordHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("senha nao bate com do hash")
	}
	return err
}
