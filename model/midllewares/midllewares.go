package midllewares

import (
	"ec2/model/loggers"
	"ec2/model/token"
	"errors"
	"net/http"
)

func Authentic(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := token.Tokenvalid(r); err != true {
			loggers.ResponseErrors(w, http.StatusInternalServerError, errors.New("token nao é valido"))
			return
		}

		next(w, r)
	}
}
