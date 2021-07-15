package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ReamTp/camposol-backend/db"
	"github.com/ReamTp/camposol-backend/models"
)

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var departament models.Departments

	// Convertir a JSON
	err := json.NewDecoder(r.Body).Decode(&departament)

	if err != nil {
		http.Error(w, "Error al recibir datos "+err.Error(), http.StatusBadRequest)
	}

	// Hacer validaciones
	if len(departament.Nombre) == 0 {
		http.Error(w, "Se requiere de un nombre para el departamento", http.StatusBadRequest)
	}

	// Buscar Departamento
	_, encontrado, _ := db.CheckDepartmentExists(departament.Nombre, false)

	if encontrado {
		http.Error(w, "Error Nombre encontrado en la base de datos", http.StatusFound)
		return
	}

	// Insertar departamento
	_, status, err := db.InsertDepartment(departament)

	if err != nil {
		http.Error(w, "Ocurrio un error al tratar de insertar el nuevo Departamento "+err.Error(), http.StatusConflict)
		return
	}

	if !status {
		http.Error(w, "No se pudo crear Departamento", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
