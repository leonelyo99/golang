package configuration

import(
	"os"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"encoding/json"
	_"github.com/go-sql-driver/mysql" //_ ejecuta sin que lo llame
)

type configuration struct {
	Server string
	Port string
	User string
	Password string
	Database string
}

//va a dar la configuracon del json
func getConfiguration() configuration{
	var c configuration
	file, err := os.Open("./config.json")  //esto decodifica json

	//corroborando que la libreria se cargue bien
	if(err != nil){
		log.Fatal(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)

	//corroborando que se decodifique bien
	if err != nil{
		log.Fatal(err)
	}

	return c
}

//obtiene la conexion a la base de datos
func GetConnection() *gorm.DB {
	c := getConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User,c.Password,c.Server,c.Port, c.Database) //devuelve un string formateado

	db, err := gorm.Open("mysql", dsn)
	if err != nil{
		log.Fatal(err)
	}

	return db
}