package handlers

import (
	"encoding/json"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"

	"net/http"
	"os"
	"rest_api_with_two_endpoints/data"
	"rest_api_with_two_endpoints/records"

	"rest_api_with_two_endpoints/verification"
)

const (
	invalidJsonMessage string = "невалидный json"
)

type errMessage struct {
	Message string `json:"message"`
}

func LoginHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		setHeaders(w)
		auth := data.Auth{}
		err := json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			ResponseError(w, http.StatusBadRequest, errors.New(invalidJsonMessage))
			return
		}

		userName, err := verification.VerificationUser(auth.UserName, auth.Password)
		if err != nil {
			ResponseError(w, http.StatusBadRequest, err)
			return
		}

		tk := &data.Token{UserId: userName}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
		w.Header().Set("Authorization", "Bearer "+tokenString)
		w.WriteHeader(http.StatusOK)
	}
}

func ListRecords() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		result := records.GetRecords()
		setHeaders(w)
		json.NewEncoder(w).Encode(result)
	}
}

func CorrectsTheEntry(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	updateEntry := data.NewEntry{}
	setHeaders(w)
	err := json.NewDecoder(r.Body).Decode(&updateEntry)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, errors.New(invalidJsonMessage))
		return
	}

	result, err := records.CorrectEntry(updateEntry.RecordNumber, updateEntry.Entry)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}
