package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TypeUsers struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo       string             `bson:"titulo" json:"titulo"`
	Descripcion  string             `bson:"descripcion" json:"descripcion,omitempty"`
	CreationDate time.Time          `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate   time.Time          `bson:"updateDate" json:"updateDate,omitempty"`
}
