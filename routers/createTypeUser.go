package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ReamTp/camposol-backend/db"
	"github.com/ReamTp/camposol-backend/models"
)

// CreateTypeUser funcion para crear Tipos de Usuarios
func CreateTypeUser(w http.ResponseWriter, r *http.Request) {
	var typeUser models.TypeUsers

	// Convertir a JSON
	err := json.NewDecoder(r.Body).Decode(&typeUser)

	if err != nil {
		http.Error(w, "Error al recibir datos "+err.Error(), http.StatusBadRequest)
	}

	// Hacer validaciones
	if len(typeUser.Titulo) == 0 {
		http.Error(w, "El Titulo es requerido", http.StatusBadRequest)
	}

	// Buscar Tipo de Usuario
	_, encontrado, _ := db.CheckTypeUserExists(typeUser.Titulo, false)

	if encontrado {
		http.Error(w, "Ya existe un tipo de usuario con ese titulo", http.StatusBadRequest)
		return
	}

	// Insertar Tipo de Usuario
	_, status, err := db.InsertTypeUser(typeUser)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), http.StatusConflict)
		return
	}

	if !status {
		http.Error(w, "No se logro registrar el tipo de Usuario", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
