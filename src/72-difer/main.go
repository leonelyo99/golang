package main
import (
	"fmt"
)

//cuando se termine de ejecutar todo se ejecuta defer por mas que lo de arriba falle
//se ejecutan de manera inversa a como se ponen en el defer

func main(){
	fmt.Println("Conectando a la base de datos")
	defer cerrarDB()

	fmt.Println("Consultamos informacion, set de datos")
	defer cerrarSetDeDatos()
}

func cerrarDB(){
	fmt.Println("Cerrar la base de datos")
}

func cerrarSetDeDatos(){
	fmt.Println("Cerrar el set de datos")
}