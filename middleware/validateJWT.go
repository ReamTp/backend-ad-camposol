package middleware

import (
	"net/http"

	"github.com/ReamTp/camposol-backend/routers"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener Authorization del header el cual contiene el token
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error en el token! "+err.Error(), http.StatusConflict)
			return
		}

		next.ServeHTTP(w, r)
	}
}
