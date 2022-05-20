package token

import (
	congif "ec2/model/config"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id int64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 12).Unix()
	permissoes["id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	tokens, err := token.SignedString([]byte(congif.SECRET))
	if err != nil {
		return "", nil
	}

	return tokens, nil
}

func Tokenvalid(r *http.Request) bool {
	tokenstr := veriToken(r)
	token, err := jwt.Parse(tokenstr, compareToken)
	if err != nil {
		return false
	}

	return token.Valid
}

func TokenIdUser(r *http.Request) (int64, error) {
	tokenstr := veriToken(r)
	token, err := jwt.Parse(tokenstr, compareToken)
	if err != nil {
		return 0, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.ParseInt(fmt.Sprintf("%.0f", permissoes["id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, err
}

func veriToken(r *http.Request) string {
	token := r.Header.Get("authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func compareToken(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return []byte(congif.SECRET), nil
	}

	return []byte(congif.SECRET), nil
}
