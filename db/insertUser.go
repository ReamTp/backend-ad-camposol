package db

import (
	"context"
	"time"

	"github.com/ReamTp/camposol-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertUser funcion para crear un nuevo usuario
func InsertUser(u models.Users) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("camposol")
	col := db.Collection("users")

	u.CreationDate = time.Now()
	u.UpdateDate = time.Now()

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
