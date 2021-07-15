package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ReamTp/camposol-backend/db"
	"github.com/ReamTp/camposol-backend/jwt"
	"github.com/ReamTp/camposol-backend/models"
)

// Login verifica el inicio de session de un usuario
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalida "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Correo) == 0 {
		http.Error(w, "El correo es requerido", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "La password debe tener al menos 8 caracteres", http.StatusBadRequest)
		return
	}

	documento, existe := db.Login(t.Correo, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o contraseña Invalidas", http.StatusNotFound)
		return
	}

	jwtKey, err := jwt.GenerateJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente "+err.Error(), http.StatusConflict)
		return
	}

	resp := models.LoginAnswer{
		Token: jwtKey,
	}

	// Seteando Contenido
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(time.Hour * 168)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
