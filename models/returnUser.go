package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnUser estructura para devolver informacion de usuarios en la base de datos
type ReturnUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre    string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos string             `bson:"apellidos" json:"apellidos,omitempty"`
	Correo    string             `bson:"correo, omitempty" json:"correo"`
	Dni       int                `bson:"dni" json:"dni,omitempty"`
	Celular   int                `bson:"celular" json:"celular,omitempty"`
	Telefono  int                `bson:"telefono" json:"telefono,omitempty"`
	Tipo      struct {
		ID     string `bson:"_id" json:"id,omitempty"`
		Titulo string `bson:"titulo" json:"titulo,omitempty"`
	}
	Departamento struct {
		ID     string `bson:"_id" json:"id,omitempty"`
		Nombre string `bson:"nombre" json:"nombre,omitempty"`
	}
	Sueldo          float64   `bson:"sueldo" json:"sueldo,omitempty"`
	FechaNacimiento time.Time `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
}
