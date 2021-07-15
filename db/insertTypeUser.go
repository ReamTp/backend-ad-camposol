package db

import (
	"context"
	"time"

	"github.com/ReamTp/camposol-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTypeUser funcion para crear un nuevo tipo de usuario
func InsertTypeUser(tu models.TypeUsers) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("camposol")
	col := db.Collection("typeUsers")

	tu.CreationDate = time.Now()
	tu.UpdateDate = time.Now()

	result, err := col.InsertOne(ctx, tu)

	if err != nil {
		return "", false, err
	}

	Obj_id := result.InsertedID.(primitive.ObjectID)

	return Obj_id.String(), true, nil
}
