package handlers

import (
	"net/http"

	"github.com/Gidraff/task-manager-webapp/utils"
)

// Error page

// Index page
func Index(w http.ResponseWriter, r *http.Request) {
	t := utils.ParseTemplateFiles("layout", "public.navbar", "index")
	t.Execute(w, nil)
}
