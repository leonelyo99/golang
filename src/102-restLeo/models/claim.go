package models

import (
	"github.com/dgrijalva/jwt-go"
)

//token de usuario
type Claim struct {
	User `json:"user"`//se save que es de tipo user
	jwt.StandardClaims //datos de un token
}