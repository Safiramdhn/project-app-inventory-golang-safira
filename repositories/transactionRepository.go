package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	// "strconv"
	"strings"

	// "fmt"
	// "strconv"
	// "strings"
	// "time"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
)

type TransactionRepositoryDB struct {
	DB *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{DB: db}
}

func (repo *TransactionRepositoryDB) GetByID(id int) (*models.Transaction, error) {
	var transaction models.Transaction
	sqlStatement := `SELECT t.id, i.id, i.name, i.price, quantity, total_price, timestamp FROM transactions t
	JOIN items i ON t.item_id = i.id
	WHERE id = $1`
	err := repo.DB.QueryRow(sqlStatement, id).Scan(&transaction.ID, &transaction.Item.ID, &transaction.Item.Name, transaction.Item.Price&transaction.Quantity, &transaction.TotalPrice, &transaction.Timestamp)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repo *TransactionRepositoryDB) GetAll(limit, offset int) ([]models.Transaction, error) {
	// Query to get the total count of all transactions
	countSqlStatement := `SELECT count(*) FROM transactions t JOIN items i ON t.item_id = i.id`
	var totalCount int
	err := repo.DB.QueryRow(countSqlStatement).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	var sqlStatement string
	values := []interface{}{}

	// Build main query with or without LIMIT/OFFSET based on the limit value
	if limit > 0 {
		sqlStatement = `SELECT t.id, t.type, i.id, i.name, i.price, t.quantity, t.total_price, t.timestamp, t.description
		                FROM transactions t JOIN items i ON t.item_id = i.id ORDER BY t.created_at ASC LIMIT $1 OFFSET $2`
		values = append(values, limit, offset)
	} else {
		sqlStatement = `SELECT t.id, t.type, i.id, i.name, i.price, t.quantity, t.total_price, t.timestamp, t.description
		                FROM transactions t JOIN items i ON t.item_id = i.id ORDER BY t.created_at ASC`
	}

	rows, err := repo.DB.Query(sqlStatement, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		transaction.Pagination.CountData = totalCount
		transaction.Pagination.Page = offset
		transaction.Pagination.PerPage = limit
		err := rows.Scan(&transaction.ID, &transaction.TransactionType, &transaction.Item.ID, &transaction.Item.Name, &transaction.Item.Price, &transaction.Quantity, &transaction.TotalPrice, &transaction.Timestamp, &transaction.Description)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (repo *TransactionRepositoryDB) Add(transaction *models.Transaction) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	sqlStatement := `SELECT price, quantity FROM items WHERE id = $1`
	err = repo.DB.QueryRow(sqlStatement, transaction.Item.ID).Scan(&transaction.Item.Price, &transaction.Item.Quantity)
	if err == sql.ErrNoRows {
		return errors.New("item not found")
	} else if err != nil {
		return err
	}

	if transaction.Quantity > transaction.Item.Quantity {
		return errors.New("insufficient quantity")
	} else {
		transaction.Item.Quantity -= transaction.Quantity
		transaction.TotalPrice = transaction.Item.Price * transaction.Quantity
		sqlStatement = `UPDATE items SET quantity = $1 WHERE id = $2`
		_, err = tx.Exec(sqlStatement, transaction.Item.Quantity, transaction.Item.ID)
	}

	sqlStatement = `INSERT INTO transactions (item_id, type, added_by, quantity, total_price, description, timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(sqlStatement, transaction.Item.ID, transaction.TransactionType, transaction.AddedBy, transaction.Quantity, transaction.TotalPrice, transaction.Description, transaction.Timestamp)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (repo *TransactionRepositoryDB) GetAllWithFilter(filter models.Transaction, limit, offset int) ([]models.Transaction, error) {
	fields := make(map[string]interface{})

	// Build filters based on the provided filter fields
	if filter.AddedBy != 0 {
		fields["added_by"] = filter.AddedBy
	}

	if filter.TransactionType != "" {
		fields["type"] = filter.TransactionType
	}

	if filter.Item.Name != "" {
		fields["name"] = "%" + filter.Item.Name + "%"
	}

	whereClauses := []string{}
	values := []interface{}{}
	index := 1
	for field, value := range fields {
		if field == "name" {
			whereClauses = append(whereClauses, fmt.Sprintf("%s LIKE $%d", field, index))
		} else {
			whereClauses = append(whereClauses, fmt.Sprintf("%s = $%d", field, index))
		}
		values = append(values, value)
		index++
	}

	// If no filters are specified, handle the error gracefully
	if len(whereClauses) == 0 {
		return nil, errors.New("no fields to filter")
	}

	// Query to get the total count of items that match the filters
	countSqlStatement := fmt.Sprintf(
		"SELECT count(*) FROM transactions t JOIN items i ON t.item_id = i.id WHERE %s",
		strings.Join(whereClauses, " AND "),
	)
	var totalCount int
	err := repo.DB.QueryRow(countSqlStatement, values...).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	// Main query to get paginated results
	var sqlStatement string
	// index = 1 + len(fields) // Start index for the values array is the length of the fields map plus 1
	if limit > 0 {
		sqlStatement = fmt.Sprintf(
			"SELECT t.id, t.type, i.id, i.name, i.price, t.quantity, t.total_price, t.timestamp, t.description "+
				"FROM transactions t JOIN items i ON t.item_id = i.id WHERE %s ORDER BY t.created_at ASC LIMIT $%d OFFSET $%d",
			strings.Join(whereClauses, " AND "),
			index,
			index+1,
		)
		values = append(values, limit, offset)
	} else {
		sqlStatement = fmt.Sprintf(
			"SELECT t.id, t.type, i.id, i.name, i.price, t.quantity, t.total_price, t.timestamp, t.description "+
				"FROM transactions t JOIN items i ON t.item_id = i.id WHERE %s ORDER BY t.created_at ASC LIMIT $%d OFFSET $%d",
			strings.Join(whereClauses, " AND "),
			index,
			index+1,
		)
		values = append(values, limit, offset)
	}

	rows, err := repo.DB.Query(sqlStatement, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		transaction.Pagination.Page = filter.Pagination.Page
		transaction.Pagination.PerPage = filter.Pagination.PerPage
		transaction.Pagination.CountData = totalCount // Set the total count here
		err := rows.Scan(&transaction.ID, &transaction.TransactionType, &transaction.Item.ID, &transaction.Item.Name, &transaction.Item.Price, &transaction.Quantity, &transaction.TotalPrice, &transaction.Timestamp, &transaction.Description)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// func (repo *TransactionRepositoryDB) Update(transaction *models.Transaction) error {
// 	tx, err := repo.DB.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p)
// 		} else if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	sqlStatement := `SELECT price, quantity FROM items WHERE id = $1`
// 	err = repo.DB.QueryRow(sqlStatement, transaction.Item.ID).Scan(&transaction.Item.Price, &transaction.Item.Quantity)

// 	fields := make(map[string]interface{})

// 	if transaction.Item.ID != 0 {
// 		fields["item_id"] = transaction.Item.ID
// 	}

// 	if transaction.TransactionType != "" {
// 		fields["type"] = transaction.TransactionType
// 	}

// 	if transaction.AddedBy != 0 {
// 		fields["added_by"] = transaction.AddedBy
// 	}

// 	if transaction.Quantity != 0 {
// 		fields["quantity"] = transaction.Quantity
// 		sqlStatement = `UPDATE items SET quantity = quantity - $2 WHERE id = $1`
// 		_, err = tx.Exec(sqlStatement, transaction.Item.ID, transaction.Quantity)
// 	}

// 	if transaction.Description != "" {
// 		fields["description"] = transaction.Description
// 	}

// 	fields["updated_at"] = time.Now()
// 	fields["total_price"] = transaction.Quantity * transaction.Item.Price
// 	setClauses := []string{}
// 	values := []interface{}{}
// 	index := 1
// 	for field, value := range fields {
// 		setClauses = append(setClauses, field+"=$"+strconv.Itoa(index))
// 		values = append(values, value)
// 		index++
// 	}

// 	if len(setClauses) == 0 {
// 		return errors.New("no fields to update")
// 	}

// 	sqlStatement = fmt.Sprintf("UPDATE transactions SET %s WHERE id = $%d AND status = 'active'", strings.Join(setClauses, ", "), index)
// 	_, err = tx.Exec(sqlStatement, values...)
// 	if err != nil {
// 		return err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (repo *TransactionRepositoryDB) Delete(id int) error {
// 	tx, err := repo.DB.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p)
// 		} else if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	sqlStatement := `UPDATE transactions SET status = 'deleted' WHERE id = $1`
// 	_, err = repo.DB.Exec(sqlStatement, id)
// 	if err != nil {
// 		return err
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return err
// 	}
// 	return nil
// }
