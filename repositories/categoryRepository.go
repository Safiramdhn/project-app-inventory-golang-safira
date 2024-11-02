package repositories

import (
	"database/sql"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
)

type CategoryRepositoryDB struct {
	DB *sql.DB
}

func NewCategoryRepositoryDB(db *sql.DB) *CategoryRepositoryDB {
	return &CategoryRepositoryDB{DB: db}
}

func (repo *CategoryRepositoryDB) Add(category *models.Category) error {
	sqlStatement := `INSERT INTO categories (name) VALUES ($1)`
	_, err := repo.DB.Exec(sqlStatement, category.Name)
	if err != nil {
		return err
	}
	return nil
}
