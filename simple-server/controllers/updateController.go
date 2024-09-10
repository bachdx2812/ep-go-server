package controllers

import (
	"api-server/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func UpdateInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := parseUpdateInfoRequest(r, &user); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if !user.IsExisted() {
		http.Error(w, "Invalid user", http.StatusNotFound)
		return
	}

	if err := user.UpdateInfo(); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User info updated!"))
}

func parseUpdateInfoRequest(r *http.Request, user *models.User) error {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Error parsing request:", err)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	info := r.FormValue("info")

	if username == "" || password == "" {
		return fmt.Errorf("username and password are required")
	}

	// Extract the file (file input field should be named 'file')
	file, handler, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("error retrieving the file: %v", err)
	}
	defer file.Close()

	// Optionally, you can save the file to a directory
	fileName := time.Now().Format("2006-01-02_15.04.05") + handler.Filename
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", fileName))
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	user.Username = username
	user.Password = password
	user.Info = info
	user.ProfileImage = fileName

	return nil
}
