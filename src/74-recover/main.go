package main
import (
	"fmt"
)

func main(){
	f()
}

func f(){

	defer func(){
		//recover trata de salvar las papas cuando hay un panico
		//son instrucciones de respaldo por si hay errores
		if r := recover(); r != nil {
			fmt.Printf("%T\n",r)
			fmt.Println("Recuperado en f: ",r)
		}
	}()

	fmt.Println("llamando a g,")
	g(4)
	fmt.Println("Retorna normalmente desde g")
}

func g(i int){
	if i>3 {
		fmt.Println("entro en panico")
		panic("El numero no puede ser mayor que 3")
	}
	defer fmt.Println("Defer en la funcion g")
	fmt.Println("Imprimiendp el valor de g ",i)
}