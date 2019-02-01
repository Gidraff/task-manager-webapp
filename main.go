package main

import (
	"log"
	"net/http"

	"github.com/Gidraff/task-manager-webapp/handlers"
	"github.com/Gidraff/task-manager-webapp/utils"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/index", handlers.Index)
	mux.HandleFunc("/signup", handlers.Signup)
	mux.HandleFunc("/signup_account", handlers.SignUpAccount)
	mux.HandleFunc("/authenticate", handlers.Authenticate)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
