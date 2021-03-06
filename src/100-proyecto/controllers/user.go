package controllers

import (
	"log"
	"fmt"
	"net/http"
	"strconv"
	"encoding/base64"
	"encoding/json"
	"crypto/sha256"
	"crypto/md5"
	"github.com/gorilla/mux"
	"102-restLeo/configuration"
	"102-restLeo/commons"
	"102-restLeo/models"
)

//Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request){
	user := models.User{}//instancio el modelo
	err := json.NewDecoder(r.Body).Decode(&user) //recivo el usuario y lo instancio en modelo

	if err != nil {
		fmt.Fprintf(w,"Error: %s\n",err)
		return
	}

	db := configuration.GetConnection()//creo la conexion a la bd
	defer db.Close()

	//cifro la contraseña que mando el usuario en sha256
	c := sha256.Sum256([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32]) //devuelve el base 16

	// busco el usuairo en la bd y comparo que el haash guardado en la bd sea el mismo que metio el usuario
	db.Where("email = ? and password = ?", user.Email, pwd).First(&user)

	if user.ID > 0 {//si es mayor a cero es que lo encontro
		user.Password = ""
		user.ConfirmPassword = ""
		token := commons.GenerateJWT(user) //genero el token
		j,err := json.Marshal(models.Token{Token:token}) //creo la respuesta en un json

		if err != nil{
			log.Fatalf("Error al convertir el token a json: %s", err)
		}
		
		//lo devuelvo con un status 200
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		//si sale mal devvuelvo el status de error
		m := models.Message{
			Message: "Usuario o clave no valido",
			Code: http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}

}
	//para registrar un usuario
func UserCreate(w http.ResponseWriter, r *http.Request){
	user := models.User{}//instancio el modelo del usuario
	m := models.Message{}//instancio el modelo del mensaje

	err := json.NewDecoder(r.Body).Decode(&user)//paso el ussuaio del request a mi usuario

	if err != nil{
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar: %s",err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w,m)
		return
	}

	if user.Password != user.ConfirmPassword{
		m.Message = "Las contraseñas no coinciden"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	//cifro la contraseña para guardarlo
	c:= sha256.Sum256([]byte(user.Password))
	pwd := base64.URLEncoding.EncodeToString(c[:32])
	user.Password = pwd

	//busco la img
	picmd5 := md5.Sum([]byte(user.Email)) //codifico el email en md5 para usar grabatar
	picstr := fmt.Sprintf("%x", picmd5) //esta funcion me devuelve un string
	user.Picture = "https://gravatar.com/avatar/"+picstr+"?s=100"

	//abro la coneccion
	db := configuration.GetConnection()//creo la conexion a la bd
	defer db.Close()
	
	//creo el usuario
	err = db.Create(&user).Error

	if err != nil{
		m.Message = fmt.Sprintf("Error al crear ek registro: %s",err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w,m)
		return
	}

	m.Message="usuario creado con exito"
	m.Code = http.StatusCreated
	commons.DisplayMessage(w, m)

}

func UserGet(w http.ResponseWriter, r *http.Request){
	user := models.User{}//instancio el modelo
	m := models.Message{}//instancio el modelo del mensaje
	
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil{
		m.Message = "Debe enviar un id valido"
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w,m)
		return
	}

	db := configuration.GetConnection()//creo la conexion a la bd
	defer db.Close()

	// busco el usuairo en la bd y comparo que el haash guardado en la bd sea el mismo que metio el usuario
	db.Where("id = ?", id).First(&user)

	if user.ID > 0 {//si es mayor a cero es que lo encontro
		user.Password = ""
		user.ConfirmPassword = ""

		u,err := json.Marshal(user) //creo la respuesta en un json

		if err != nil{
			log.Fatalf("Error al convertir el token a json: %s", err)
		}
		
		//lo devuelvo con un status 200
		w.WriteHeader(http.StatusOK)
		w.Write(u)
	} else {
		//si sale mal devvuelvo el status de error
		m := models.Message{
			Message: "Usuario no encontrado",
			Code: http.StatusBadRequest,
		}
		commons.DisplayMessage(w, m)
	}

}