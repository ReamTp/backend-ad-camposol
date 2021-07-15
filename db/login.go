package db

import (
	"github.com/ReamTp/camposol-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.ReturnUser, bool) {
	var userResult models.ReturnUser
	use, find, _ := CheckUserExists(email)
	depart, findD, _ := CheckDepartmentExists(use.Departamento, true)
	typeU, findTU, _ := CheckTypeUserExists(use.Tipo, true)

	if !find || !findD || !findTU {
		return userResult, false
	}

	// Convertir contraseñas
	passwordBytes := []byte(password)
	passwordDB := []byte(use.Password)

	// Validar contraseñas
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return userResult, false
	}

	userResult.ID = use.ID
	userResult.Nombre = use.Nombre
	userResult.Apellidos = use.Apellidos
	userResult.Correo = use.Correo
	userResult.Dni = use.Dni
	userResult.Celular = use.Celular
	userResult.Telefono = use.Telefono
	userResult.Tipo.ID = typeU.ID.Hex()
	userResult.Tipo.Titulo = typeU.Titulo
	userResult.Departamento.ID = depart.ID.Hex()
	userResult.Departamento.Nombre = depart.Nombre
	userResult.Sueldo = use.Sueldo
	userResult.FechaNacimiento = use.FechaNacimiento

	return userResult, true
}
