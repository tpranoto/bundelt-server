package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/tpranoto/bundelt-server/common"
	"github.com/tpranoto/bundelt-server/internal/app/users"
	"github.com/tpranoto/bundelt-server/internal/storage"
)

var usersApp users.App
var logger *log.Logger

func init() {
	router := mux.NewRouter()
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.SetOutput(os.Stdout)

	usersApp = users.App{
		Router:      router,
		Port:        ":" + common.GetEnv("PORT", "13000"),
		UserStorage: storage.NewPostgreSQLStorage(logger),
		Logger:      logger,
	}
}

func main() {
	logger.Println("Starting bundelt server...")

	usersApp.Run()
}
