package controllers

import (
	"api-server/models"
	"errors"
	"fmt"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	if err := parseSignUpRequest(r, &user); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if user.IsExisted() {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	if err := user.Register(); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func parseSignUpRequest(r *http.Request, user *models.User) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	// Retrieve form values
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		return errors.New("username and password are required")
	}

	user.Username = username
	user.Password = password

	fmt.Printf("Username: %s, Password: %s\n", user.Username, user.Password)
	return nil
}
