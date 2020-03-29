package main
import (
	"net/http"
	"log"
)

func main(){

	//===============================================
	//Manera normal
	//===============================================

	// //si te pegan a esa url, anda a esa carpeta
	// http.Handle("/", http.FileServer(http.Dir("public")))
	// log.Println("ejecutando server en http;//localhost:8080")
	// //en donde lo voy a servir ej: localhost, el log para imprimir que pasa en pantallla
	// log.Println(http.ListenAndServe(":8080", nil))

	//===============================================
	//con multiplexor (recive varias solicitudes al tiempo y las procesa)
	//===============================================
	mux := http.NewServeMux()
	//si te pegan a esa url, anda a esa carpeta
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	//en donde lo voy a servir ej: localhost, el log para imprimir que pasa en pantallla
	log.Println("ejecutando server en http;//localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}