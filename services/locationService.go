package services

import (
	"database/sql"
	"errors"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/repositories"
)

func CreateLocation(db *sql.DB, location models.Location) error {
	if location.Name == "" {
		return errors.New("location name cannot be empty")
	}
	if location.Address == "" {
		return errors.New("location address cannot be empty")
	}

	LocationRepo := repositories.NewLocationRepositoryDB(db)
	return LocationRepo.Add(&location)
}
