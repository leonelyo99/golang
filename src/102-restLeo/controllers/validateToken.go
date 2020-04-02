package controllers

import(
	"log"
	"context" //enviar informacion de usuario entre los controladores
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"102-restLeo/commons"
	"102-restLeo/models"
)

//ValidateToken validar el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m models.Message
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)

	log.Println("El token es")
	log.Println(token)
	log.Println("El error es")
	log.Println(err)
	//si el proceso de validar un token genera un error
	if err != nil{
		m.Code = http.StatusUnauthorized
		switch  err.(type) {
		case *jwt.ValidationError:
			vErr :=err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				m.Message="Su token ha expirado"
				commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message="Token a sido modificado"
				commons.DisplayMessage(w, m)
				return
			default:
				m.Message="Su token no es valido"
				commons.DisplayMessage(w, m)
				return
			}
		}
	}

	if token.Valid{
		//obtener la informacion y devolverla en el requestc
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		log.Println("contenido del token")
		log.Println(ctx)
		next(w,r.WithContext(ctx))
	}else{
		m.Code = http.StatusUnauthorized
		m.Message="Su token no es valido"
				commons.DisplayMessage(w, m)
				return
	}
}