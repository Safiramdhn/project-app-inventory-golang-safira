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
	case "search_item":
	case "view_all_items":
	case "get_item_by_id":
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
