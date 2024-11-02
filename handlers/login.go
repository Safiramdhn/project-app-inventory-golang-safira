package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/services"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/utils"
)

func LoginHandler(db *sql.DB) {
	var user_input models.User
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

		if err := decoder.Decode(&user_input); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	userFound, err := services.Login(db, user_input)
	if err != nil {
		errMessage := fmt.Sprintf("Error while logging in: %v", err)
		response = models.Response{StatusCode: 400, Message: errMessage, Data: nil}
	} else {
		response = models.Response{StatusCode: 200, Message: "Logged in successfully", Data: userFound}
		// generate session
		go utils.CreateSession(userFound)

	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))
}

func LogoutHandler() {
	err := utils.DeleteSession()
	if err != nil {
		fmt.Println("Error deleting session: ", err)
	}
	fmt.Println("Logged out successfully")
	response := models.Response{StatusCode: 200, Message: "Logged out successfully", Data: nil}
	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON: ", err)
	}
	fmt.Println(string(responseJson))
	os.Exit(0)
}
