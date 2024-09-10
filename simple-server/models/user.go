package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ProfileImage string `json:"profileImage"`
	Info         string `json:"info"`
}

func (u *User) IsExisted() bool {
	users, err := readUsersFromFile()
	if err != nil {
		return false
	}

	for _, user := range users {
		if user.Username == u.Username {
			return true
		}
	}
	return false
}

func (u *User) Register() error {
	users, err := readUsersFromFile()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Username == u.Username {
			return fmt.Errorf("user already exists")
		}
	}

	users = append(users, *u)

	jsonData, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %v", err)
	}

	// Write the JSON string to a file
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %v", err)
	}

	return nil
}

func (u *User) UpdateInfo() error {
	users, err := readUsersFromFile()
	if err != nil {
		return err
	}

	for i, user := range users {
		if user.Username == u.Username {
			users[i] = *u
			break
		}
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("could not marshal JSON: %v", err)
	}

	// Write the JSON string to a file
	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %v", err)
	}

	return nil
}

// Function to read the data.json file and return the slice of users
func readUsersFromFile() ([]User, error) {
	// Open the JSON file
	file, err := os.Open("data.json")
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	// Unmarshal the JSON data into a slice of User structs
	var users []User
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	return users, nil
}
