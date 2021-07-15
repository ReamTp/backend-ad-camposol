package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users estructura para registrar nuevos usuarios en la base de datos
type Users struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	Correo          string             `bson:"correo, omitempty" json:"correo"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Dni             int                `bson:"dni" json:"dni,omitempty"`
	Celular         int                `bson:"celular" json:"celular,omitempty"`
	Telefono        int                `bson:"telefono" json:"telefono,omitempty"`
	Tipo            string             `bson:"tipo" json:"tipo,omitempty"`
	Departamento    string             `bson:"departamento" json:"departamento,omitempty"`
	Sueldo          float64            `bson:"sueldo" json:"sueldo,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Activo          bool               `bson:"activo" json:"activo,omitempty"`
	CreationDate    time.Time          `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate      time.Time          `bson:"updateDate" json:"updateDate,omitempty"`
}
