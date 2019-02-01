package handlers

import (
	"net/http"

	"github.com/Gidraff/task-manager-webapp/data"
	"github.com/Gidraff/task-manager-webapp/utils"
)

// Login handler
// Show the login page
func Login(writer http.ResponseWriter, request *http.Request) {
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// Signup handler
// Show the signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	utils.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// SignUpAccount handler
func SignUpAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		utils.Danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		utils.Danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// Authenticate user
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	// check for errors
	if err != nil {
		utils.Danger(err, "Cannot find user")
	}
	// compare passwords
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			utils.Danger(err, "Cannot create session")
		}
		// create a session cookie
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		// set cookie
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
