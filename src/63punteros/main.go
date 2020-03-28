package main
import "fmt"

func main(){

	a := 3
	//==================================
	//punteros
	//==================================

	fmt.Println("Antes de duplicar, a =",a);
	//paso el valor en memoria del contenedor de a
	duplicar(&a)
	fmt.Println("Antes de duplicar, a =",a);
}

func duplicar(x *int){
	//tomo el contenedor de ese espacio en memoria
	*x = *x * 2
	fmt.Println("Dentro de la funcion duplicar x vale", *x)
}