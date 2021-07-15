package db

import (
	"context"
	"time"

	"github.com/ReamTp/camposol-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists recibe un email para buscar si existe en la DB
func CheckUserExists(email string) (models.Users, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("camposol")
	col := db.Collection("users")

	// Formato json
	condicion := bson.M{"correo": email}

	var resultado models.Users

	// Buscar elemento
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	// Convertir objectId a hexadecimal
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
