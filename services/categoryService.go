package services

import (
	"database/sql"
	"errors"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/repositories"
)

func CreateCategory(db *sql.DB, category models.Category) error {
	if category.Name == "" {
		return errors.New("category name cannot be empty")
	}

	categoryRepo := repositories.NewCategoryRepositoryDB(db)
	return categoryRepo.Add(&category)
}
