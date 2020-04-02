package commons

import (
	"log"
	"crypto/rsa"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
	"102-restLeo/models"
)

//forma de declarar variables de paquetes
var (
	privateKey *rsa.PrivateKey //esta llave almacena la llave privada

	//se usa oara validar el token
	PublicKey *rsa.PublicKey //esta llave almacena la llave publica
)

func init(){
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa") //lee un archivo del sistema operativo
	log.Println("leyendo el private key")
	if err != nil{
		log.Fatal("No se pudo leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil{
		log.Fatal("No se pudo leer el archivo publico")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil{
		log.Fatal("No se pudo hacer el parse a privateKey")
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil{
		log.Fatal("No se pudo hacer el parse a PublicKey")
	}
}

//Genera el token para el cliente
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour*2).Unix(), // para expirar la sesion
			Issuer: "Escuela Digital",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil{
		log.Fatal("No se pudo firmar el token")
	}

	return result
}