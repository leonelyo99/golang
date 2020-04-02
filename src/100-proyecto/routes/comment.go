package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"102-restLeo/controllers"
)

func SetCommentRouter(router *mux.Router){
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true) //uso obligatorio del slash despues del prefix
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST") //esto comprobara el token

	router.PathPrefix(prefix).Handler(
		negroni.New(
			// negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}