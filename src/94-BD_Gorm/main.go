package main
import (
	"fmt"
	_"github.com/go-sql-driver/mysql" //_ ejecuta solo el main del paquete
	"github.com/jinzhu/gorm"
)

//modelo
type Producto struct{
	gorm.Model //ID, CreteAt, UpadateAt, DeletedAt
	CodigoBarras string
	Precio uint
}

func main(){
	db, err := gorm.Open("mysql", "root:sasa@/edcurso?charset=utf8&parseTime=True&loc=Local")

	//si hay error en la conexion avisa
	if err != nil{
		panic("Error en la conexion a la base de datos")
	}

	//cierra la base de datos al final
	defer db.Close()
	fmt.Println("Se conecto con la base datos")

	//crea la tabla productos
	// db.CreateTable(&Producto{})
	// fmt.Println("Se crea la tabla productos en la base de datos")

	//insertar elemeneto en la bd
	// db.Create(&Producto{CodigoBarras: "dasdad123123123", Precio: 12200})

	//busca el elemento y lo almacena en esa variable
	var p Producto
	db.First(&p)
	fmt.Println(p)

	//edito el producto
	p.Precio = 400
	db.Save(&p)
	fmt.Println("El nuevo precio del producto es: ",p.Precio)
}
