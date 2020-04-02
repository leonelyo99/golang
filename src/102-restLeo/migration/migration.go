package migration

import(
	"102-restLeo/configuration"
	"102-restLeo/models"
)

//permite crear las tablas en la base de datos
func Migrate(){
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Perfil{})
}