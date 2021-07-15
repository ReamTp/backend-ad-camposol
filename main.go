package main

import (
	"log"

	"github.com/ReamTp/camposol-backend/controllers"
	"github.com/ReamTp/camposol-backend/db"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("Sin conexion a la DB")
		return
	}

	controllers.Controladores()
}
