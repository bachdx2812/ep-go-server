package controllers

import (
	"api-server/models"
	"fmt"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := parseSignInRequest(r, &user); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if !user.IsExisted() {
		http.Error(w, "User does not exist / Password is incorrect", http.StatusNotFound)
		return
	}

	w.Write([]byte("User logged in"))
}

func parseSignInRequest(r *http.Request, user *models.User) error {
	// Parse the form (for form-data)
	err := r.ParseForm()
	if err != nil {
		return err
	}

	// Retrieve form values
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Username: %s, Password: %s\n", username, password)

	// Validate that username and password are not empty
	if username == "" || password == "" {
		return fmt.Errorf("username and password are required")
	}

	user.Username = username
	user.Password = password

	return nil
}
