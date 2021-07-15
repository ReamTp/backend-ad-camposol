package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ReamTp/camposol-backend/db"
	"github.com/ReamTp/camposol-backend/models"
)

// CreateUser funcion para crear en la base de datos un usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.Users

	// Convertir a JSON
	err := json.NewDecoder(r.Body).Decode(&userModel)

	if err != nil {
		http.Error(w, "Error al recibir los datos "+err.Error(), http.StatusBadRequest)
	}

	// Hacer validaciones
	if len(userModel.Correo) == 0 {
		http.Error(w, "El correo es requerido", http.StatusBadRequest)
	}

	if len(userModel.Password) < 8 {
		http.Error(w, "La contraseña debe tener como minimo 8 caracteres", http.StatusBadRequest)
	}

	// Buscar Usuario
	_, encontrado, _ := db.CheckUserExists(userModel.Correo)

	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email", http.StatusConflict)
		return
	}

	// Insertar Usuario
	_, status, err := db.InsertUser(userModel)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), http.StatusConflict)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
