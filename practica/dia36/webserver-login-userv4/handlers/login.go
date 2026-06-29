package handlers

import (
	"fmt"
	"html/template"
	"main/config"
	"main/pkg"
	"main/repository"
	"main/service"
	"net/http"

	"github.com/gorilla/securecookie"
)

const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
	<button type="submit">Logout</button>
</form>
`

func InternalTmplHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		fmt.Fprintf(w, internalPage, userName)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}

	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func LoginTmplHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pass := r.FormValue("password")
		redirectTarget := "/"

		repo := repository.NewSQLUserRepository(config.GetConnection())
		service := service.NewUserService(repo)

		user, err := service.GetUserByEmail(email)
		if err != nil {
			http.Error(w, "No se encontro usuario", http.StatusBadRequest)
			return
		}

		matchPassword := pkg.CheckPasswordHash(pass, user.Password)

		if !matchPassword {
			http.Error(w, "Password incorrecto", http.StatusBadRequest)
			return
		}

		setSession(email, w)
		redirectTarget = "/internal"

		http.Redirect(w, r, redirectTarget, 302)
	}

	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

}

func LogoutTmplHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/login", 302)
}
