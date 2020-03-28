package animales
import "fmt"

type Gato struct{
	Nombre string
}

func (p Gato) Comunicarse(){
	fmt.Println("Hola soy un gato")
}