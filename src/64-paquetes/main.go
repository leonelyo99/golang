package main
import (
	"64-paquetes/saludar"
	"fmt"
)
func main(){

	saludar.Nombre = "leonel"

	fmt.Println(saludar.Nombre)

	//==================================
	//paquetes
	//==================================
	saludar.Saludar("leonel")
}