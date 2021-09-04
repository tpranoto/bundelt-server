package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tpranoto/bundelt-server/internal/storage"
)

func (a *App) handlerUserGroupAdd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	input := storage.UserGroupRel{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	err = a.UserGroupRelStorage.WriteUserGroupRelation(input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handlerUserGroupDetailGet(w http.ResponseWriter, r *http.Request) {
	paramUser := r.FormValue("user")

	resDt, err := a.UserGroupRelStorage.FindGroupsByUserFbId(paramUser)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	groups, err := a.GroupStorage.GetGroupDetailByIds(resDt)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(groups)
	w.Write(res)
}

func (a *App) handlerDeleteUserGroup(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	input := storage.UserGroupRel{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	err = a.UserGroupRelStorage.DeleteUserGroup(input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handlerDeleteUserGroupDetails(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	input := storage.UserGroupRel{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	err = a.UserGroupRelStorage.DeleteUserGroup(input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	err = a.GroupStorage.DeleteGroupDetails(input.GroupId)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handlerAddUserGroupDetails(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	input := storage.UserGroupDetailInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	input1 := storage.UserGroupRel{
		UserFbId: input.UserFbId,
		GroupId:  input.GroupId,
	}

	lat, _ := strconv.ParseFloat(input.Latitude, 64)
	lon, _ := strconv.ParseFloat(input.Longitude, 64)

	input2 := storage.GroupInfo{
		GroupID:   input.GroupId,
		GroupName: input.GroupName,
		Desc:      input.Desc,
		Created:   input.Created,
		Latitude:  lat,
		Longitude: lon,
	}

	err = a.UserGroupRelStorage.WriteUserGroupRelation(input1)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	err = a.GroupStorage.WriteGroupDetails(input2)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
