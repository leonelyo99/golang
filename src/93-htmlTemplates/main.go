package main
import (
	"net/http"
	"html/template"
	"log"
)

type Persona struct {
	Nombre string
	Edad uint8
}

//enviar datos a un template
func rederizar( w http.ResponseWriter, r *http.Request){
	p:= &Persona{"Daniel",26}
	
	//t = template, err= errror
	t, err := template.ParseFiles("./views/index.html")
	
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, p)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", rederizar)

	log.Println("ejecutando server en http;//localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
	log.Println("Ya no esta corriendo el servidor")
}