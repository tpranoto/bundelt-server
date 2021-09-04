package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tpranoto/bundelt-server/common/middleware"
	"github.com/tpranoto/bundelt-server/internal/storage"
)

type App struct {
	Router              *mux.Router
	Port                string
	UserStorage         storage.UsersStorage
	UserGroupRelStorage storage.UserGroupsStorage
	GroupStorage        storage.GroupStorage
	Logger              *log.Logger
}

func (a *App) assignRoutes() {
	a.httpHandler(http.MethodGet, "/", middleware.MultiMiddlwares(a.handlerHealthCheck))
	a.httpHandler(http.MethodGet, "/user", middleware.MultiMiddlwares(a.handlerGetUser))
	a.httpHandler(http.MethodPost, "/signup", middleware.MultiMiddlwares(a.handlerSignUp))
	a.httpHandler(http.MethodPost, "/login", middleware.MultiMiddlwares(a.handlerHealthCheck))

	//user groups
	a.httpHandler(http.MethodPost, "/user_group/add", middleware.MultiMiddlwares(a.handlerUserGroupAdd))
	a.httpHandler(http.MethodPost, "/user_group/delete", middleware.MultiMiddlwares(a.handlerDeleteUserGroup))
	a.httpHandler(http.MethodGet, "/user_group/get", middleware.MultiMiddlwares(a.handlerUserGroupDetailGet))
	a.httpHandler(http.MethodPost, "/user_group_detail/add", middleware.MultiMiddlwares(a.handlerAddUserGroupDetails))
	a.httpHandler(http.MethodPost, "/user_group_detail/delete", middleware.MultiMiddlwares(a.handlerDeleteUserGroupDetails))

	//groups
	a.httpHandler(http.MethodPost, "/group/add", middleware.MultiMiddlwares(a.handleGroupAdd))
	a.httpHandler(http.MethodPost, "/group/delete", middleware.MultiMiddlwares(a.handleDeleteGroupDetail))
	a.httpHandler(http.MethodGet, "/group/nearby", middleware.MultiMiddlwares(a.handleNearbyGroups))
}

func (a *App) httpHandler(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(pattern, handler).Methods(method)
}

func (a *App) Run() {
	//assign all http routes
	a.assignRoutes()

	a.Logger.Printf("Starting serve bundelt users on port %s", a.Port)
	err := http.ListenAndServe(a.Port, a.Router)
	if err != nil {
		a.Logger.Fatalf("Failed to serve bundelt users: %v", err)
	}
}
