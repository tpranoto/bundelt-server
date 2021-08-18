package users

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	App struct {
		Router *mux.Router
		Port   string
		Logger *log.Logger
	}
)

func (a *App) Run() {
	//assign all http routes
	a.assignRoutes()

	a.Logger.Printf("Serving bundelt users on port %s", a.Port)
	err := http.ListenAndServe(a.Port, a.Router)
	if err != nil {
		a.Logger.Fatalf("Failed to serve bundelt users: %v", err)
	}
}
