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
	sqlStatement := `SELECT id, name, quantity, price, category_id, location_id FROM items WHERE id = $1 AND status = 'active'`
	row := repo.DB.QueryRow(sqlStatement, id)

	err := row.Scan(&item.ID, &item.Name, &item.Quantity, &item.Price, &item.CategoryId, &item.LocationId)
	if err == sql.ErrNoRows {
		return nil, errors.New("item not found")
	} else if err != nil {
		return nil, err
	}

	return &item, nil
}

func (repo *ItemRepositoryDB) GetAll(limit, offset int) ([]models.Item, error) {
	countSqlStatement := `SELECT count(*) FROM items`
	var totalCount int
	err := repo.DB.QueryRow(countSqlStatement).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	var items []models.Item
	var sqlStatement string
	values := []interface{}{}

	if limit > 0 {
		sqlStatement = `SELECT id, name, quantity, price, category_id, location_id FROM items WHERE status = 'active' ORDER BY created_at ASC LIMIT $1 OFFSET $2`
		values = append(values, limit, offset)
	} else {
		sqlStatement = `SELECT id, name, quantity, price, category_id, location_id FROM items WHERE status = 'active'`
	}

	rows, err := repo.DB.Query(sqlStatement, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		item.Pagination.CountData = totalCount
		item.Pagination.Page = offset
		item.Pagination.PerPage = limit

		err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Price, &item.CategoryId, &item.LocationId)
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

func (repo *ItemRepositoryDB) GetAllWithFilter(item models.Item, limit, offset int) ([]models.Item, error) {
	fields := make(map[string]interface{})

	// Build filters based on the provided item fields
	if item.Name != "" {
		fields["name"] = "%" + item.Name + "%"
	}
	if item.Quantity != 0 {
		fields["quantity"] = item.Quantity
	}
	if item.Price != 0 {
		fields["price"] = item.Price
	}
	fields["status"] = "active" // Ensure we only get active items

	whereClauses := []string{}
	values := []interface{}{}
	index := 1

	for field, value := range fields {
		if field == "name" {
			// Use LIKE for partial match on "name" field
			whereClauses = append(whereClauses, fmt.Sprintf("%s LIKE $%d", field, index))
		} else {
			whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", field, index))
		}
		values = append(values, value)
		index++
	}

	countSqlStatement := fmt.Sprintf(
		"SELECT count(*) FROM items WHERE %s",
		strings.Join(whereClauses, " AND "),
	)
	var totalCount int
	err := repo.DB.QueryRow(countSqlStatement, values...).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	var sqlStatement string
	// Construct the SQL query with pagination
	if limit > 0 {
		sqlStatement = fmt.Sprintf(
			"SELECT id, name, quantity, price, category_id, location_id FROM items WHERE %s ORDER BY created_at ASC LIMIT $%d OFFSET $%d",
			strings.Join(whereClauses, " AND "),
			index,
			index+1,
		)
		values = append(values, limit, offset)

	} else {
		sqlStatement = fmt.Sprintf(
			"SELECT id, name, quantity, price, category_id, location_id FROM items WHERE %s",
			strings.Join(whereClauses, " AND "),
		)
	}

	// Add limit and offset to the values

	rows, err := repo.DB.Query(sqlStatement, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Price, &item.CategoryId, &item.LocationId)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
