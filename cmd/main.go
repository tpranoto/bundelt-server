package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"

	"github.com/tpranoto/bundelt-server/internal/app/users"
)

var usersApp users.App
var logger *log.Logger

func init() {
	router := mux.NewRouter()
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.SetOutput(os.Stdout)

	usersApp = users.App{
		Router: router,
		Port:   ":" + os.Getenv("PORT"),
		Logger: logger,
	}
}

func main() {
	logger.Println("Starting bundelt server...")

	usersApp.Run()
}
