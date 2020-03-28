package main
import "fmt"

// Persona es una estructura
type Persona struct {
	nombre string;
	edad uint8;
}

func main(){
	//=============================================
	//Structs
	//=============================================
	var persona1 Persona
	persona1.edad = 23
	persona1.nombre = "Leonel"
	fmt.Println(persona1);

	persona2 :=Persona{nombre:"Juan", edad:23}
	fmt.Println(persona2);
}
