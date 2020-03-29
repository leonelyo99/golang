package main
import (
	"fmt"
)

func main(){
	bodegaOrigen := []string{"php", "javascript","mysql","mongoDb","java","python"}
	bodegaDestino := []string{}

	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	go func(){
		for _ , libro := range bodegaOrigen{
			fotocopiadora <- libro
		}
	}()

	//se carga la fotocopiadora
	go func(){
		for {
			//entrega el contenido de la fotocopiadora a la variable libro
			//descargando el canal <- izquierda
			libro := <-fotocopiadora
			//cargando el canal  <- derecha
			fotocopiado <- libro

			//agregando el libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino){
				//cerrar el canal fotocopiado
				close(fotocopiado)
			}
		}
	}()

	fmt.Println("**Listado de libros fotocopiados**")
	for {
		//el ok comprueba el canal de fotocopiado
		if libro, ok := <- fotocopiado; ok{
			fmt.Println(libro)
		}else{
			break
		}
	}
}