package controllers

import(
	"net/http"
	// "fmt"
	// "encoding/json"
	// "102-restLeo/models"
	// "102-restLeo/configuration"
	// "102-restLeo/commons"
)

func CommentCreate(w http.ResponseWriter, r *http.Request){
	// comment := models.Comment{}
	// m := models.Message{}

	// err:= json.NewDecoder(r.Body).Decode(&comment) //comprobamos si el usuario envio bien los parametros
	
	// if err != nil{
	// 	m.Code =http.StatusBadRequest
	// 	m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
	// 	commons.DisplayMessage(w, m)
	// 	return
	// }

	// db := configuration.GetConnection()
	// defer db.Close()

	// err = db.Create(&comment).Error

	// if err != nil{
	// 	m.Code =http.StatusBadRequest
	// 	m.Message = "Error al crear el comentario"
	// 	commons.DisplayMessage(w, m)
	// 	return
	// }

	// m.Code = http.StatusCreated
	// m.Message="Creado con exito"
	// commons.DisplayMessage(w,m)
}