package main
import "fmt"

func main(){
	//=============================================
	//Slices
	//=============================================
	//Array de tamaño variable
	//var nombres[]string
	
	//make(tipo de dato, tamaño inicial, capacidad inicial)
	nombres := make([]string,0)

	nombres = append(nombres, "Daniel")
	nombres = append(nombres, "Leonel")
	nombres = append(nombres, "Lautaro")

	fmt.Println(nombres);

}
