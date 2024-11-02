package repositories

import (
	"database/sql"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
)

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{DB: db}
}

func (repo *UserRepositoryDB) Login(user models.User) (*models.User, error) {
	// var userFound models.User
	sqlStatement := `SELECT id, role FROM users WHERE username = $1 AND password = $2`

	// Execute the SQL statement
	err := repo.DB.QueryRow(sqlStatement, user.Username, user.Password).Scan(&user.ID, &user.Role)

	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	user.Password = "" // Hide password from response
	return &user, nil
}
