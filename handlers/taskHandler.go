package handlers

import (
	"net/http"

	"github.com/Gidraff/task-manager-webapp/utils"
)

// Dashboard returns dashboard template
func Dashboard(writer http.ResponseWriter, request *http.Request) {
	_, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	}
	t := utils.ParseTemplateFiles("layout", "dashboard", "side.navbar", "private.navbar")
	t.Execute(writer, nil)
}

// NewTask handler
func NewTask(writer http.ResponseWriter, request *http.Request) {
	_, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	}
	t := utils.ParseTemplateFiles("layout", "task", "side.navbar", "private.navbar")
	t.Execute(writer, nil)
}

// CreateTask creates
func CreateTask(writer http.ResponseWriter, request *http.Request) {
	sess, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			utils.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			utils.Danger(err, "Cannot get user from session")
		}
		task := request.PostFormValue("task")
		if _, err := user.CreateTask(task); err != nil {
			utils.Danger(err, "Cannot create task")
		}
		http.Redirect(writer, request, "/task", 302)
	}
}
