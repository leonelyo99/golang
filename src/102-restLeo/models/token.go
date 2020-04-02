package models

//token permite envolver el token generado (claim)
type Token struct {
	Token string `json:"token"`
}