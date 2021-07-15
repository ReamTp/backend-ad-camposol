package routers

import (
	"errors"
	"strings"

	"github.com/ReamTp/camposol-backend/db"
	"github.com/ReamTp/camposol-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("REAM")

	// Modelo de chequeo debe ser un punturo
	claims := &models.Claim{}

	// Separar token de la palabra Bearer
	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	// Obtener token
	tk = strings.TrimSpace(splitToken[1])

	// Procesar Token
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := db.CheckUserExists(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}

	return claims, false, "", err
}
