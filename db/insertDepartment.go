package db

import (
	"context"
	"time"

	"github.com/ReamTp/camposol-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertDepartment Funcion para crear un nuevo departamento
func InsertDepartment(d models.Departments) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("camposol")
	col := db.Collection("departments")

	d.CreationDate = time.Now()
	d.UpdateDate = time.Now()

	result, err := col.InsertOne(ctx, d)

	if err != nil {
		return "", false, err
	}

	Obj_id := result.InsertedID.(primitive.ObjectID)

	return Obj_id.String(), true, nil
}
