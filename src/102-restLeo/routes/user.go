package routes

import(
	"github.com/gorilla/mux" //enrutador
	"github.com/urfave/negroni" //middleware
	"102-restLeo/controllers"
)

//SetUserRouter ruta para el registro de usuario
func SetUserRouter(router *mux.Router){
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")
	subRouter.HandleFunc("/{id}", controllers.UserGet).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}