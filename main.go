package main

import (
	"fmt"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/configs"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/handlers"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
)

func main() {
	var user *models.User
	db, err := configs.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// fmt.Println("Database connected successfully!")

	user = handlers.LoginHandler(db)

	if user == nil {
		return
	}

	switch user.Role {
	case "admin":
		fmt.Println("Welcome admin!")
	case "staff_warehouse":
		fmt.Println("Welcome user!")
	default:
		fmt.Println("Invalid role!")
	}
}
