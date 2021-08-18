package users

import (
	"net/http"

	"github.com/tpranoto/bundelt-server/common/middleware"
)

func (a *App) assignRoutes() {
	a.httpHandler(http.MethodGet, "/", middleware.MultiMiddlwares(handlerHealthCheck))
}

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello World`))
}

func (a *App) httpHandler(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	a.Router.HandleFunc(pattern, handler).Methods(method)
}
