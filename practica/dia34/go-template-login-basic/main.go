package main

import (
	"main/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.LoginPage)
	http.HandleFunc("/login", handlers.SingupPage)
	http.HandleFunc("/welcome", handlers.WelcomePage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8085", nil)
}
