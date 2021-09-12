package http

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tpranoto/bundelt-server/common"
	"github.com/tpranoto/bundelt-server/internal/storage"
)

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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.UserInfoInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = a.UserStorage.CreateUserInfo(input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handlerLogIn(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.UserLoginInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userInfo, loginFailed, err := a.UserStorage.GetUserLogin(input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.UserStorage.UpdateUserInfo(storage.UpdateUserInfoInput{
		UserID: userInfo.UserID,
		Latitude: common.UpdateFloat64{
			Update: true,
			Value:  input.Lat,
		},
		Longitude: common.UpdateFloat64{
			Update: true,
			Value:  input.Lon,
		},
	})
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if loginFailed {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "_SID_Bundelt",
		Value:    base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", userInfo.UserID))),
		Path:     "/",
		Domain:   a.Domain,
		Expires:  time.Now().AddDate(0, 0, 60),
		MaxAge:   5184000, // 60 days
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	res, _ := json.Marshal(userInfo)
	w.Write(res)
}

func (a *App) handlerLogOut(w http.ResponseWriter, r *http.Request) {

}
