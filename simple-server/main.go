package main

import (
	"api-server/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/signUp", controllers.SignUpHandler)
	http.HandleFunc("/signIn", controllers.SignInHandler)
	http.HandleFunc("/updateInfo", controllers.UpdateInfoHandler)
	http.ListenAndServe(":8080", nil)
}
