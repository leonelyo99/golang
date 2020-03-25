package main
import "fmt"

func main(){

	//=====================================================
	//variables
	//=====================================================

	//se declara la variable y el tipo despues del nombre,
	//se pueden declarar dos variables en el mismo comando
	var nombre, apellido string
	nombre="leonel"
	apellido="Mazzan"

	//se pueden autodeclarar las variables
	edad:=20
	//se pueden autodeclarar las variables juntas
	altura,peso := 1.80, 80

	fmt.Println(nombre+" "+apellido)
	fmt.Println(edad,peso,altura)

	/*=====================================================
	constantes
	=====================================================*/
	const perro = "Akira" 
	fmt.Println(perro)
}