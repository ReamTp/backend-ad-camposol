package middleware

import (
	"net/http"

	"github.com/ReamTp/camposol-backend/db"
)

// CheckDB es el middleware que permite revisar el estado de la db
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Conexion con la base de datos perdida", 500)
		}

		// Pasar los valores recibidos a la funcion http.HandlerFunc que recibimos
		next.ServeHTTP(w, r)
	}
}
