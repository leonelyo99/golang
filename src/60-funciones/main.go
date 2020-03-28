package main
import "fmt"

func main(){

	//==================================
	//funcion
	//==================================
	saludar("leonel","mazzan")
	
	//retornar un dato
	fmt.Println(suma(2,3))

	// retornar multiples datos
	n :=[]int8{32,43,6,54,74,22,2}
	maximo, minimo := maxMin(n)
	fmt.Println(maximo, minimo)

	//==================================
	//funcion anonima
	//==================================
	anonima := func (){
		fmt.Println("me inprimo desde una funcion anonima")
	}
	anonima()
}

func saludar(nombre string, apellido string){
	fmt.Println("hola "+nombre+" "+apellido)
}

func suma(a, b int) int{
	return a + b
}

//retornar multiples valores
func maxMin(numeros []int8)(int8, int8){
	var max, min int8

	for _ , v := range numeros{
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max, min
}