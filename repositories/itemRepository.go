package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
)

type ItemRepositoryDB struct {
	DB *sql.DB
}

func NewItemRepositoryDB(db *sql.DB) *ItemRepositoryDB {
	return &ItemRepositoryDB{DB: db}
}

func (repo *ItemRepositoryDB) GetByID(id int) (*models.Item, error) {
	var item models.Item
	sqlStatement := `SELECT id, name, quantity, price FROM items WHERE id = $1`
	row := repo.DB.QueryRow(sqlStatement, id)

	err := row.Scan(&item.ID, &item.Name, &item.Quantity, item.Price)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &item, nil
}

func (repo *ItemRepositoryDB) GetAll() ([]models.Item, error) {
	var items []models.Item
	sqlStatement := `SELECT id, name, quantity, price FROM items`
	rows, err := repo.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (repo *ItemRepositoryDB) Add(item *models.Item) error {
	sqlStatement := `INSERT INTO items (name, category_id, location_id, quantity, price) VALUES ($1, $2, $3, $4, $5)`
	_, err := repo.DB.Exec(sqlStatement, item.Name, item.CategoryId, item.LocationId, item.Quantity, item.Price)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ItemRepositoryDB) Update(item *models.Item) error {
	fields := make(map[string]interface{})

	if item.Name != "" {
		fields["name"] = item.Name
	}
	if item.Quantity != 0 {
		fields["quantity"] = item.Quantity
	}
	if item.Price != 0 {
		fields["price"] = item.Price
	}
	if item.CategoryId != 0 {
		fields["catergory_id"] = item.CategoryId
	}
	if item.LocationId != 0 {
		fields["location_id"] = item.LocationId
	}

	fields["updated_at"] = time.Now()
	setClauses := []string{}
	values := []interface{}{}
	index := 1
	for field, value := range fields {
		setClauses = append(setClauses, field+"=$"+strconv.Itoa(index))
		values = append(values, value)
		index++
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	sqlStatement := fmt.Sprintf("UPDATE items SET %s WHERE id = $%d AND status = 'active'", strings.Join(setClauses, ", "), index)
	values = append(values, item.ID)
	_, err := repo.DB.Exec(sqlStatement, values...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ItemRepositoryDB) Delete(id int) error {
	sqlStatement := `UPDATE items SET status = 'deleted' WHERE id = $1`
	_, err := repo.DB.Exec(sqlStatement, id)
	return err
}

func (repo *ItemRepositoryDB) Search(item models.Item) ([]models.Item, error) {
	fields := make(map[string]interface{})

	if item.Name != "" {
		fields["name"] = item.Name
	}
	if item.Quantity != 0 {
		fields["quantity"] = item.Quantity
	}
	if item.Price != 0 {
		fields["price"] = item.Price
	}

	whereClauses := []string{}
	values := []interface{}{}
	index := 1
	for field, value := range fields {
		if field == "name" {
			// Use LIKE for the "name" field
			whereClauses = append(whereClauses, field+" LIKE $"+strconv.Itoa(index))
		} else {
			whereClauses = append(whereClauses, field+" = $"+strconv.Itoa(index))
		}
		values = append(values, value)
		index++
	}

	if len(whereClauses) == 0 {
		return nil, errors.New("no fields to update")
	}

	sqlStatement := fmt.Sprintf("SELECT id, name, quantity, price FROM items WHERE %s AND status = 'active'", strings.Join(whereClauses, " AND "))
	rows, err := repo.DB.Query(sqlStatement, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
