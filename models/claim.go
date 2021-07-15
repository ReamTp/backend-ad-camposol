package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claim estructura para procesar JWT
type Claim struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email string             `json:"email"`
	jwt.StandardClaims
}
