package main
import "fmt"

func main(){
	//==================================================
	//if
	//==================================================
	if (5<10) {
		fmt.Println("Esto esta en el bloque de verdadero")
	}
	
	isValid := true

	//admite variables
	if (isValid){
		fmt.Println("Esto esta en el bloque de verdadero")
	}

	//admite instrucciones
	if isValid2 := true; isValid2{
		fmt.Println("Esto esta en el bloque de verdadero")
	}

	//==================================================
	//swich
	//==================================================
	
	casoA := 5;

	switch casoA{ //no admite instrucciones como en el if. 
		case 4: //se pueden poner condiciones
			fmt.Println("caso 1") //no hace falta el break;
		case 2:
			fmt.Println("caso 2")
			fallthrough//si pongo esto sigue evaluando por mas que se cumpla la condicion
		default:
			fmt.Println("caso por defecto")
	}
}
