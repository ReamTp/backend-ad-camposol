package db

import (
	"context"
	"time"

	"github.com/ReamTp/camposol-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckDepartmentExists(valor string, id bool) (models.Departments, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("camposol")
	col := db.Collection("departments")

	var resultado models.Departments
	var condicion primitive.M

	if id {
		idHex, _ := primitive.ObjectIDFromHex(valor)
		condicion = bson.M{"_id": idHex}
	} else {
		condicion = bson.M{"nombre": valor}
	}

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
