package routes

import(
	"github.com/gorilla/mux" //enrutador
)

//funcion para inicia las rutas
func InitRoutes() *mux.Router{
	router := mux.NewRouter().StrictSlash(false)
	SetUserRouter(router)
	SetLoginRouter(router)
	SetCommentRouter(router)

	return router
}