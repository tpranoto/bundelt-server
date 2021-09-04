package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tpranoto/bundelt-server/common"
	"github.com/tpranoto/bundelt-server/internal/storage"
)

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
		return
	}

	input := storage.CreateUserInfoInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	err = a.UserStorage.CreateUserInfo(input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
