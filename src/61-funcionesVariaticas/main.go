package main
import "fmt"

func main(){

	//==================================
	//funcion variatica
	//==================================
	saludarVarios(23, "juan","pepe","roberto","dario","daniel")
}

//recive un numero variante de parametros y solo puede tener un parametro variante
func saludarVarios(edad uint8, nombres ...string){ //adentro recivira un slice
	fmt.Printf("%T\n",nombres)

	for _, v:= range nombres {
		fmt.Println("Hola",v, "tienes",edad)
	}
}

