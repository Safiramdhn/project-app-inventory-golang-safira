package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/services"
)

func GetAllTransactionsHandler(db *sql.DB) {
	var transactions []models.Transaction
	var pagination models.Pagination
	var response models.Response
	repoDB := struct {
		DB *sql.DB
	}{
		DB: db,
	}

	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&pagination); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
		}
	}
	fmt.Printf("per page: %d, page: %d\n", pagination.PerPage, pagination.Page)

	transactionsService := services.NewTransaction(repoDB)
	transactions, err = transactionsService.GetAllTransactions(pagination)
	if err != nil {
		errMessage := fmt.Sprintf("Error getting items: %v", err)
		response = models.Response{StatusCode: 500, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Items retrieved successfully", Data: transactions}
	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))
}

func GetTransactionByIdHandler(db *sql.DB) {
	var transaction *models.Transaction
	var response models.Response
	repoDB := struct {
		DB *sql.DB
	}{
		DB: db,
	}

	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&transaction); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
		}
	}

	transactionsService := services.NewTransaction(repoDB)
	transaction, err = transactionsService.GetTransactionByID(transaction.ID)
	if err != nil {
		errMessage := fmt.Sprintf("Error getting item with id %d: %v", transaction.ID, err)
		response = models.Response{StatusCode: 500, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Item retrieved successfully", Data: transaction}
	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}

	fmt.Println(string(responseJson))
}

func CreateTransactionHandler(db *sql.DB, id int) {
	var transaction models.Transaction
	var response models.Response
	repoDB := struct {
		DB *sql.DB
	}{
		DB: db,
	}

	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&transaction); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
		}
	}

	transaction.AddedBy = id
	transactionsService := services.NewTransaction(repoDB)
	err = transactionsService.CreateTransaction(&transaction)
	if err != nil {
		errMessage := fmt.Sprintf("Error creating item: %v", err)
		response = models.Response{StatusCode: 500, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 201, Message: "Item created successfully", Data: nil}
	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))
}

func GetTransactionsWithFilterHandler(db *sql.DB) {
	var transaction models.Transaction
	var response models.Response
	repoDB := struct {
		DB *sql.DB
	}{
		DB: db,
	}

	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&transaction); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
		}
	}

	TransactionsService := services.NewTransaction(repoDB)
	transactions, err := TransactionsService.GetAllTransactionsWithFilter(transaction)
	if err != nil {
		errMessage := fmt.Sprintf("Error getting Transactions: %v", err)
		response = models.Response{StatusCode: 500, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Transactions retrieved successfully", Data: transactions}
	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))
}

// func UpdateTransactionHandler(db *sql.DB) {
// 	var transaction models.Transaction
// 	var response models.Response
// 	repoDB := struct {
// 		DB *sql.DB
// 	}{
// 		DB: db,
// 	}

// 	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
// 	if err != nil {
// 		fmt.Println("Open file error message: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Check if the file is empty
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		fmt.Println("File stat error: ", err)
// 		return
// 	}

// 	if fileInfo.Size() > 0 {
// 		decoder := json.NewDecoder(file)

// 		if err := decoder.Decode(&transaction); err != nil && err != io.EOF {
// 			fmt.Println("Decode error message: ", err)
// 		}
// 	}

// 	transactionsService := services.NewTransaction(repoDB)
// 	err = transactionsService.UpdateTransaction(&transaction)
// 	if err != nil {
// 		errMessage := fmt.Sprintf("Error updating item with id %d: %v", transaction.ID, err)
// 		response = models.Response{StatusCode: 500, Message: errMessage, Data: nil}
// 	} else {
// 		response = models.Response{StatusCode: 200, Message: "Item updated successfully", Data: nil}
// 	}

// 	responseJson, err := json.MarshalIndent(response, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON: ", err)
// 	}
// 	fmt.Println(string(responseJson))
// }

// func DeleteTransactionHandler(db *sql.DB) {
// 	var transaction models.Transaction
// 	var response models.Response
// 	repoDB := struct {
// 		DB *sql.DB
// 	}{
// 		DB: db,
// 	}

// 	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
// 	if err != nil {
// 		fmt.Println("Open file error message: ", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Check if the file is empty
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		fmt.Println("File stat error: ", err)
// 		return
// 	}

// 	if fileInfo.Size() > 0 {
// 		decoder := json.NewDecoder(file)

// 		if err := decoder.Decode(&transaction); err != nil && err != io.EOF {
// 			fmt.Println("Decode error message: ", err)
// 		}
// 	}

// 	transactionsService := services.NewTransaction(repoDB)
// 	err = transactionsService.DeleteTransaction(transaction.ID)
// 	if err != nil {
// 		errMessage := fmt.Sprintf("Error deleting item with id %d: %v", transaction.ID, err)
// 		response = models.Response{StatusCode: 500, Message: errMessage, Data: nil}
// 	} else {
// 		response = models.Response{StatusCode: 200, Message: "Item deleted successfully", Data: nil}
// 	}

// 	responseJson, err := json.MarshalIndent(response, "", "  ")
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON: ", err)
// 		return
// 	}
// 	fmt.Println(string(responseJson))
// }
