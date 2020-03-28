package main
import "fmt"

func main(){

	//==================================
	//for
	//==================================
	for i := 0; i<5; i++ {
		if(i ==2){
			break
		}
		fmt.Println(i)
	}

	//==================================
	//for continuo (while)
	//==================================
	a:=5
	for a > 0{
		a--
		fmt.Println(a)
	}

	//==================================
	//for forever
	//==================================
	// for {
	// 	fmt.Println("por siempre")
	// }
	
	//==================================
	//foreach
	//==================================

	numeros := []string{"cero","uno","dos","tres"}

	//primero el index y despues el valor, si el index no lo uso pongo _
	for _ , i := range numeros{
		fmt.Println(i)
	}

}
