package main

import (
	"fmt"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/configs"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/handlers"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/utils"
)

func main() {
	var endpoint string
	db, err := configs.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// fmt.Println("Database connected successfully!")
	fmt.Println("endpoint: ")
	fmt.Scan(&endpoint)

	if endpoint == "login" {
		handlers.LoginHandler(db)
	} else if endpoint == "logout" {
		handlers.LogoutHandler()
	} else {
		user := utils.GetSession()
		if user == nil {
			return
		} else if user.Role == "admin" {
			handlers.AdminHandler(db, endpoint)
		} else if user.Role == "staff_warehouse" {
			handlers.StaffHandler(db, endpoint, user.ID)
		}
	}
}
