package main
import "fmt"

func main(){
	//=============================================
	//Map
	//=============================================

	idiomas := make(map[string]string);
	idiomas["es"]="Español"
	idiomas["en"]="Ingles"
	idiomas["fr"]="Frances"
	fmt.Println(idiomas)

	nombres := map[string]string{"l":"leonel","p":"pablo","j":"juan"}
	fmt.Println(nombres)

	//borrar de un map
	delete(nombres,"l")
	fmt.Println(nombres)

	//se pueden hacer comprobaciones mientras se asigna un valor, devolvera false si no lo tiene
	if nombre, ok := nombres["l"]; ok{
		fmt.Println("existe l ",nombre)
	}else{
		fmt.Println("no existe l",nombre)
	}
}
