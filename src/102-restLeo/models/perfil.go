package models

//usuario del sistema
type Perfil struct{
	ID uint `gorm:"primary_key"`
	Perfil string `json: "perfil" gorm:"not null;"`
}