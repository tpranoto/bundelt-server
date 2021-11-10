package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/tpranoto/bundelt-server/common"
	http "github.com/tpranoto/bundelt-server/internal/app"
	"github.com/tpranoto/bundelt-server/internal/storage"
)

var app http.App
var logger *log.Logger

func init() {
	router := mux.NewRouter()
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.SetOutput(os.Stdout)

	postgresDB := storage.NewPostgreSQLStorage(logger)

	app = http.App{
		Domain:              common.GetEnv("DOMAIN", "localhost"),
		Router:              router,
		Port:                ":" + common.GetEnv("PORT", "13000"),
		UserStorage:         postgresDB,
		UserGroupRelStorage: postgresDB,
		GroupStorage:        postgresDB,
		GroupMessageStorage: postgresDB,
		EventStorage:        postgresDB,
		EventGroupStorage:   postgresDB,
		UserEventStorage:    postgresDB,
		Logger:              logger,
	}
}

func main() {
	logger.Println("Starting bundelt server...")

	app.Run()
}
