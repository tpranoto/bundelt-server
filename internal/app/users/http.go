package users

import (
	"net/http"

	"github.com/tpranoto/bundelt-server/common/middleware"
)

func (a *App) assignRoutes() {
	a.httpHandler(http.MethodGet, "/", middleware.MultiMiddlwares(a.handlerHealthCheck))
	a.httpHandler(http.MethodGet, "/user", middleware.MultiMiddlwares(a.handlerGetUser))
	a.httpHandler(http.MethodPost, "/signup", middleware.MultiMiddlwares(a.handlerSignUp))
	a.httpHandler(http.MethodPost, "/login", middleware.MultiMiddlwares(a.handlerHealthCheck))
}

func (a *App) httpHandler(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(pattern, handler).Methods(method)
}
