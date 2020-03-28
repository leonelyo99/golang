package main
import "fmt"

func main(){
	//=============================================
	//Arrays
	//=============================================
	var nombres [10] string;
	nombres[0]="leonel";
	fmt.Println(nombres)

	//se puede auto declarar
	apellidos := [3]string{"mazzan","perez","fernandez"}
	fmt.Println(apellidos)
	fmt.Println(len(apellidos))

	//=====================================
	//Slices
	//=====================================
	var SliceS []string = apellidos[1:3] //datos de la pocision 1 a la 4
	fmt.Println(SliceS)
}
