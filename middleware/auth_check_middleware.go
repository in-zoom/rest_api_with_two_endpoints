package middleware

import (
	"context"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"rest_api_with_two_endpoints/data"
	"rest_api_with_two_endpoints/handlers"
	"strings"
)

func AuthCheckMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		w.Header().Add("Content-Type", "application/json")
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			err := errors.New("Отсутствует токен авторизации")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			err := errors.New("Неверный токен авторизации")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		tokenPart := splitted[1]
		tk := &data.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		if err != nil {
			err := errors.New("Неверный токен аутентификации")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		if !token.Valid {
			err := errors.New("Токен недействителен")
			handlers.ResponseError(w, http.StatusForbidden, err)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next(w, r, ps)
	}
}

func UpdateEntry() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.CorrectsTheEntry(w, r, ps)
	}
}
