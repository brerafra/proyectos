package handlers

import (
	"fmt"
	"main/config"
	"main/pkg"
	"main/repository"
	"main/service"
	"net/http"

	"github.com/gorilla/securecookie"
)

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
	<label for="email">Email</label>
	<input type="text" id="email" name="email">
	<label for="password">Password</label>
	<input type="password" id="password" name="password">
	<button type="submit">Login</button>
</form>
`

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexPage)
}

const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
	<button type="submit">Logout</button>
</form>
`

func InternalPageHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		fmt.Fprintf(w, internalPage, userName)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
