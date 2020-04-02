package routes

import(
	"github.com/gorilla/mux" //enrutador
	"102-restLeo/controllers"
)

//SetLoginRouter ruta para el login de usuario
func SetLoginRouter(router *mux.Router){ //vamos a recibir una ruta que va a ser un puntero
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}