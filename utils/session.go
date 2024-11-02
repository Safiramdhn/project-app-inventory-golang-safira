package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
)

func CreateSession(user *models.User) {
	file, err := os.OpenFile("session.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	file.Seek(0, 0)  // Go back to the start of the file
	file.Truncate(0) // Clear the file content

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&user); err != nil {
		fmt.Println("Encode error message: ", err)
		return
	}
}

func GetSession() *models.User {
	file, err := os.Open("session.json")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return nil
	}
	defer file.Close()

	var user models.User
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return nil
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&user); err != nil && err != io.EOF {
			fmt.Println("Decode error message: ", err)
			return nil
		}
	} else {
		fmt.Println("There is no body data in the file")
	}

	return &user
}
