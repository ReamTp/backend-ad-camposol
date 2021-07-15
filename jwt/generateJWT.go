package jwt

import (
	"time"

	"github.com/ReamTp/camposol-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

// Genero un encriptado con JWT
func GenerateJWT(t models.ReturnUser) (string, error) {
	myKey := []byte("REAM")

	payload := jwt.MapClaims{
		"_id":       t.ID.Hex(),
		"nombre":    t.Nombre,
		"apellidos": t.Apellidos,
		"correo":    t.Correo,
		"dni":       t.Dni,
		"celular":   t.Celular,
		"telefono":  t.Telefono,
		"tipo": bson.M{
			"id":     t.Tipo.ID,
			"titulo": t.Tipo.Titulo,
		},
		"departamento": bson.M{
			"id":     t.Departamento.ID,
			"nombre": t.Departamento.Nombre,
		},
		"sueldo":           t.Sueldo,
		"fecha_nacimiento": t.FechaNacimiento,
		"exp":              time.Now().Add(time.Hour * 168).Unix(),
	}

	// Generar Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Asignar firma
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
