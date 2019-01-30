package handlers

import (
	"net/http"

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
