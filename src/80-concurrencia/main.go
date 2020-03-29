package main
import (
	"fmt"
	"time"
)

func main(){
	var h string
	//go ejecuta la funcion de manera independiente al resto del sistema(concurrencia)
	//si la funcion main termina corta la funcion que go esta ejecutando
	go MostrarNumeros()

	fmt.Print("Digita algo: ")
	fmt.Scan(&h)
	fmt.Println("Digitaste", h)
}

func MostrarNumeros(){
	for i := 1; i <= 10; i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}