package main

import (
	"fmt"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/configs"
)

func main() {
	db, err := configs.InitDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database connected successfully!")
}
