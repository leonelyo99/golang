package main
import (
	"fmt"
	"sync"
	//"runtime"// cantidad de procesadores a usar
)

func main(){
	// runtime.GOMAXPROCS(4) // cantidad de procesadores a usar
	var wg sync.WaitGroup  //estoy creando un grupo de go rutinas 

	numbers := []uint32{12332,324312,523423,32144,1321355,12312556,5474588,453352,43234,765643,1231256,7888,98,76,5,4,3,2,2}

	wg.Add(len(numbers)) //inicializando la go rutina de x cantidad de elementos

	fmt.Println("Comienza el proceso...")
	for _ , v := range numbers {
		go func(a uint32){//declaro go rutina
			defer wg.Done() //cuando termina sacala del grupo de go rutinas
			fmt.Println(a,EsPrimo(a))
		}(v)
	}
	wg.Wait() //espera que todas terminen
	fmt.Println("Termino el proceso")
}

func EsPrimo(a uint32)bool{
	c := 0
	var i uint32

	for i = 1; i <= a; i++ {

		if a%i == 0 {
			c++
		}
	}
	if c== 2{
		return true
	}
	return false
	
}