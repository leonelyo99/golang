package models

import (
	"github.com/jinzhu/gorm"
)

//usuario del sistema
type User struct{
	gorm.Model
	Username string `json: "username" gorm:"not null; unique"`
	PerfilId uint8 `json: "perfilId" gorm:"not null; unique"`
	Perfil Perfil `json: "-"`
}