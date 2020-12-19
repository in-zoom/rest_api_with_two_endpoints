package data

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type NewEntry struct {
	RecordNumber string `json:"record_number"`
	Entry        string `json:"entry"`
}

type Token struct {
	UserId string
	jwt.StandardClaims
}
