package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/urfave/negroni" //middleware
	"102-restLeo/migration"
	"102-restLeo/routes"
)

func main(){
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion") //crea la bandera
	flag.Parse()

	if migrate == "yes"{
		log.Println("Comenzo la migracion...")
		migration.Migrate()
		log.Println("finalizo la migracion...")
	}

	//inicia las rutas
	router := routes.InitRoutes()

	n:=negroni.Classic()
	n.UseHandler(router)

	//inicia el servidor
	server := &http.Server{
		Addr: ":8080",
		Handler: n,
	}

	log.Println("Iniciando el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Final de la ejecucion hasta pronto")
}