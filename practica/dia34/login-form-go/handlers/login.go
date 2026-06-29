package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func SingupPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Printf("New user signup: Username - %s, Password - %s\n", username, password)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		//perform authentication logic here (e.g. check against a database)
		//for simplicity we'll just the username and password are both admin
		if username == "admin" && password == "admin" {
			//successful logic, redirect to a welcome page.
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
			return
		}

		//Invalid credential, show the login page with an error message.
		fmt.Fprintf(w, "Invalid credentials. Please try again.")
		return
	}

	//if not a post request, serve the login page template.
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome you have successfully logged in!")
}
