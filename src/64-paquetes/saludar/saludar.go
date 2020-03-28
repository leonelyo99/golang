package saludar
import "fmt"

//con mayusculas significa que se va a exportar si no tiene mayusculas es solo local
var Nombre string

func Saludar(nombre string){
	fmt.Println("Holaa",nombre)
}