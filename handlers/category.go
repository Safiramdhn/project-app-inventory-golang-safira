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

func CreateCategoryHandler(db *sql.DB) {
	var category_input models.Category
	var response models.Response

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

		if err := decoder.Decode(&category_input); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	err = services.CreateCategory(db, category_input)
	if err != nil {
		errMessage := fmt.Sprintf("Error creating category: %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Category created successfully"}
	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))

}
