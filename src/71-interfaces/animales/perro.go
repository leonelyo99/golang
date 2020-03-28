package animales
import "fmt"

type Perro struct{
	Nombre string
}

func (p Perro) Comunicarse(){
	fmt.Println("Hola soy un perro")
}