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

func AddItemHandler(db *sql.DB) error {
	var item_input models.Item
	var response models.Response
	repoDB := struct {
		DB *sql.DB
	}{
		DB: db,
	}

	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return nil
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return nil
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&item_input); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return nil
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	itemService := services.NewItemService(repoDB)
	err = itemService.AddItem(item_input)

	if err != nil {
		errMessage := fmt.Sprintf("Error adding item: %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		errMessage := "Item added successfully"
		response = models.Response{StatusCode: 201, Message: errMessage, Data: nil}
	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))
	return nil
}

func UpdateItemHandler(db *sql.DB) error {
	var item_input models.Item
	var response models.Response
	repoDB := struct {
		DB *sql.DB
	}{
		DB: db,
	}

	file, err := os.OpenFile("body.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return nil
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return nil
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&item_input); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
		}
	}
	itemService := services.NewItemService(repoDB)
	err = itemService.UpdateItem(item_input)
	if err != nil {
		errMessage := fmt.Sprintf("Error updating item: %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		errMessage := "Item updated successfully"
		response = models.Response{StatusCode: 200, Message: errMessage, Data: nil}
	}
	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))
	return nil
}
