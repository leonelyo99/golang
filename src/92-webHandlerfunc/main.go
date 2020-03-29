package main
import (
	"fmt"
	"net/http"
	"log"
)

func messageHandlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hola este es un handlerfunc</h1>")
}

func messageHFCustom(message string) http.Handler{
	return http.HandlerFunc( 
		func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, message)
		},
	)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandlerFunc)
	mux.Handle("/saludar", messageHFCustom("<h1>Hola handlerfunc</h1>"))


	log.Println("ejecutando server en http;//localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}