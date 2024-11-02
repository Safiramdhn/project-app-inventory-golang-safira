package handlers

import (
	"database/sql"
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
		SearchItemsHandler(db)
	case "view_all_items":
		GetItemsHandler(db)
	case "get_item_by_id":
		GetItemByIDHandler(db)
	}
}

func StaffHandler(db *sql.DB, endpoint string) {
	switch endpoint {
	case "add_item":
	case "update_item":
	case "delete_item":
	case "search_item":
	case "view_all_items":
	case "get_item_by_id":
	}
}
