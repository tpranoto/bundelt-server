package users

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tpranoto/bundelt-server/common"
	"github.com/tpranoto/bundelt-server/internal/storage"
)

type App struct {
	Router      *mux.Router
	Port        string
	UserStorage storage.UsersStorage
	Logger      *log.Logger
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

func (a *App) handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello World`))
	a.Logger.Println(a.UserStorage.GetUserInfo(1))
}

func (a *App) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	paramUserID := r.FormValue("user_id")
	userID, err := common.ConvToInt64(paramUserID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userData, err := a.UserStorage.GetUserInfo(userID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(userData)
	w.Write(res)
}

func (a *App) handlerSignUp(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
	}

	input := storage.CreateUserInfoInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
	}

	err = a.UserStorage.CreateUserInfo(input)
	if err != nil {
		a.Logger.Println(err)
	}

}
