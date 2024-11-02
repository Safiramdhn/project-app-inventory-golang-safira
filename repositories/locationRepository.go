package repositories

import (
	"database/sql"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
)

type LocationRepositoryDB struct {
	DB *sql.DB
}

func NewLocationRepositoryDB(db *sql.DB) *LocationRepositoryDB {
	return &LocationRepositoryDB{DB: db}
}

func (repo *LocationRepositoryDB) Add(location *models.Location) error {
	query := `INSERT INTO locations (name, address) VALUES ($1, $2)`
	_, err := repo.DB.Exec(query, location.Name, location.Address)
	return err
}
