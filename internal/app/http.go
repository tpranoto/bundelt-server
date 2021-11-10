package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tpranoto/bundelt-server/common/middleware"
	"github.com/tpranoto/bundelt-server/internal/storage"
)

type App struct {
	Domain              string
	Router              *mux.Router
	Port                string
	UserStorage         storage.UsersStorage
	UserGroupRelStorage storage.UserGroupsStorage
	GroupStorage        storage.GroupStorage
	GroupMessageStorage storage.GroupMessageStorage
	EventStorage        storage.EventStorage
	EventGroupStorage   storage.EventGroupStorage
	UserEventStorage    storage.UserEventStorage
	Logger              *log.Logger
}

func (a *App) assignRoutes() {
	//users
	a.httpHandler(http.MethodGet, "/user", middleware.MultiMiddlwares(a.handlerGetUser))
	a.httpHandler(http.MethodPost, "/signup", middleware.MultiMiddlwares(a.handlerSignUp))
	a.httpHandler(http.MethodPost, "/login", middleware.MultiMiddlwares(a.handlerLogIn))
	a.httpHandler(http.MethodPost, "/logout", middleware.MultiMiddlwares(a.handlerLogOut))

	//user groups
	a.httpHandler(http.MethodPost, "/user_group/add", middleware.MultiMiddlwares(a.handlerUserGroupAdd))
	a.httpHandler(http.MethodPost, "/user_group/delete", middleware.MultiMiddlwares(a.handlerDeleteUserGroup))
	a.httpHandler(http.MethodGet, "/user_group/get", middleware.MultiMiddlwares(a.handlerUserGroupDetailGet))
	a.httpHandler(http.MethodGet, "/group/member", middleware.MultiMiddlwares(a.handlerGroupMemberGet))
	a.httpHandler(http.MethodPost, "/user_group_detail/add", middleware.MultiMiddlwares(a.handlerAddUserGroupDetails))
	a.httpHandler(http.MethodPost, "/user_group_detail/delete", middleware.MultiMiddlwares(a.handlerDeleteUserGroupDetails))

	//groups
	a.httpHandler(http.MethodPost, "/group/add", middleware.MultiMiddlwares(a.handleGroupAdd))
	a.httpHandler(http.MethodPost, "/group/delete", middleware.MultiMiddlwares(a.handleDeleteGroupDetail))
	a.httpHandler(http.MethodGet, "/group/nearby", middleware.MultiMiddlwares(a.handleNearbyGroups))

	//group message
	a.httpHandler(http.MethodGet, "/message/group", middleware.MultiMiddlwares(a.handleGetGroupMessage))
	a.httpHandler(http.MethodPost, "/message/group/add", middleware.MultiMiddlwares(a.handleAddNewGroupMessage))

	//event group
	a.httpHandler(http.MethodPost, "/event_group_detail/create", middleware.MultiMiddlwares(a.handlerAddEventGroupDetails))
	a.httpHandler(http.MethodPost, "/event_group/join", middleware.MultiMiddlwares(a.handlerJoinEventGroup))
	a.httpHandler(http.MethodGet, "/event/member", middleware.MultiMiddlwares(a.handlerGetEventMember))
	a.httpHandler(http.MethodGet, "/event", middleware.MultiMiddlwares(a.EventInfoHandler))
}

func (a *App) httpHandler(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(pattern, handler).Methods(method)
}

func (a *App) Run() {
	//assign all http routes
	a.assignRoutes()

	a.Logger.Printf("Starting serve bundelt on port %s", a.Port)
	err := http.ListenAndServe(a.Port, a.Router)
	if err != nil {
		a.Logger.Fatalf("Failed to serve bundelt: %v", err)
	}
}
