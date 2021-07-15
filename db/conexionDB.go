package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN objeto de la conexion a la DB
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://editData:editdata@twittor.5lwr5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ConectarDB sirve para conectar la DB
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	// Verificar si la conexion esta disponible
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión Exitosa a la DB")
	return client
}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil
}
