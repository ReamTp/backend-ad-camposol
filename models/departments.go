package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Departments struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre       string             `bson:"nombre" json:"nombre"`
	Descripcion  string             `bson:"descripcion" json:"descripcion"`
	CreationDate time.Time          `bson:"creationDate" json:"creationDate"`
	UpdateDate   time.Time          `bson:"updateDate" json:"updateDate"`
}
