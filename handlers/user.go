package handlers

import (
	"database/sql"
	"fmt"
)

func AdminHandler(db *sql.DB, endpoint string) {
	switch endpoint {
	case "add_item":
		AddItemHandler(db)
	case "update_item":
		UpdateItemHandler(db)
	case "delete_item":
		DeleteItemHandler(db)
	case "search_item":
		GetItemWithFilterHandler(db)
	case "view_all_items":
		GetItemsHandler(db)
	case "get_item_by_id":
		GetItemByIDHandler(db)
	default:
		fmt.Println("Invalid endpoint")
		return
	}
}

func StaffHandler(db *sql.DB, endpoint string, id int) {
	switch endpoint {
	case "create_transaction":
		CreateTransactionHandler(db, id)
	// case "update_transaction":
	// 	UpdateTransactionHandler(db)
	// case "delete_transaction":
	// 	DeleteTransactionHandler(db)
	case "view_transaction_history":
		GetAllTransactionsHandler(db)
	case "get_transaction_by_id":
		GetTransactionByIdHandler(db)
	case "view_all_items":
		GetItemsHandler(db)
	case "search_transaction":
		GetTransactionsWithFilterHandler(db)
	case "get_item_by_id":
	default:
		fmt.Println("Invalid endpoint")
		return
	}
}
