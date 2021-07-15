package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/ReamTp/camposol-backend/middleware"
	"github.com/ReamTp/camposol-backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Controllers inicia y pone escuchar el servidor, asi como, tambien controla las rutas
func Controladores() {
	router := mux.NewRouter()

	router.HandleFunc("/insert/user", middleware.CheckDB(routers.CreateUser)).Methods("POST")
	router.HandleFunc("/insert/tuser", middleware.CheckDB(routers.CreateTypeUser)).Methods("POST")
	router.HandleFunc("/insert/department", middleware.CheckDB(routers.CreateDepartment)).Methods("POST")

	// Setear mi puerto - buscar si existe en el sistema
	PORT := os.Getenv("PORT")

	// Comprobar no si hay puerto definido
	if PORT == "" {
		PORT = "8080"
	}

	// Definir los permisos de la api
	controller := cors.AllowAll().Handler(router)

	// Iniciar Servidor
	log.Fatal(http.ListenAndServe(":"+PORT, controller))
}
