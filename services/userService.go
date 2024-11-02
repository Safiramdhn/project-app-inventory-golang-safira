package services

import (
	"database/sql"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/repositories"
)

func Login(db *sql.DB, user models.User) (*models.User, error) {
	userRepo := repositories.NewUserRepositoryDB(db)
	return userRepo.Login(user)
}
