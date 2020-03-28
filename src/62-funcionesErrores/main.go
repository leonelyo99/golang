package main
import (
		"fmt"
		"errors"
	)

func main(){
	//==================================
	//funcion errrors
	//==================================
	// err := errors.New("mi mensaje de errorr")
	// fmt.Println(err)

	r, err := division(100,0)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}else{
		fmt.Println(r)
	}
}

func division(dividendo, divisor float64)(resultado float64, err error){
	if (divisor == 0){
		err = errors.New("No se puede dividir por cero")
	}else {
		resultado = dividendo/divisor
	}
	return
}

