package db

import (
	"context"
	"time"

	"github.com/ReamTp/camposol-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckTypeUserExists recibe un string y un bool que si es true comprueba si el mismo ya existe por ID, si no, por titulo
func CheckTypeUserExists(valor string, id bool) (models.TypeUsers, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("camposol")
	col := db.Collection("typeUsers")

	var condicion primitive.M
	if id {
		idHex, _ := primitive.ObjectIDFromHex(valor)
		condicion = bson.M{"_id": idHex}
	} else {
		condicion = bson.M{"titulo": valor}
	}

	var resultado models.TypeUsers

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
