package main
import (
	"fmt"
)

//creando una structura (clase)
type Persona struct {
	nombre string
	edad int8
}

type Numero int

//funcion de la estructura persona creada mas arriba = Metodo
func (p Persona) saludar(){
	fmt.Println("Hola soy una persona")
}

func (n Numero) presentarse(){
	fmt.Println("Hola, yo soy un numero")
}

func(p *Persona) asignarEdad(i int8){
	if(i >= 0){
		p.edad = i
	}
}

func main(){
	var persona Persona
	persona.saludar()
	persona.asignarEdad(35)
	fmt.Println(persona.edad)

	var numero Numero
	numero = 23
	fmt.Println(numero)
	numero.presentarse()

}