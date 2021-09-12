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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.UserGroupRel{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.UserGroupRelStorage.WriteUserGroupRelation(input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handlerUserGroupDetailGet(w http.ResponseWriter, r *http.Request) {
	paramUser := r.FormValue("user")

	userID, err := strconv.ParseInt(paramUser, 10, 64)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resDt, err := a.UserGroupRelStorage.FindGroupsByUserId(userID)
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.UserGroupRel{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.UserGroupRelStorage.DeleteUserGroup(input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handlerDeleteUserGroupDetails(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.UserGroupRel{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.UserGroupRelStorage.DeleteUserGroup(input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.GroupStorage.DeleteGroupDetails(input.GroupId)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handlerAddUserGroupDetails(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.UserGroupDetailInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input1 := storage.GroupInput{
		GroupName: input.GroupName,
		Desc:      input.Desc,
		Created:   input.Created,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	}

	groupInfo, err := a.GroupStorage.WriteGroupDetails(input1)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input2 := storage.UserGroupRel{
		UserId:  input.UserId,
		GroupId: groupInfo.GroupID,
	}

	err = a.UserGroupRelStorage.WriteUserGroupRelation(input2)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(storage.UserGroupDetailInfo{
		UserId:    input.UserId,
		GroupId:   groupInfo.GroupID,
		GroupName: groupInfo.GroupName,
		Desc:      groupInfo.Desc,
		Created:   groupInfo.Created,
		Latitude:  groupInfo.Latitude,
		Longitude: groupInfo.Longitude,
	})
	w.Write(res)
}

func (a *App) handlerGroupMemberGet(w http.ResponseWriter, r *http.Request) {
	paramGroupId := r.FormValue("group_id")

	groupID, err := strconv.ParseInt(paramGroupId, 10, 64)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userIDs, err := a.UserGroupRelStorage.FindMemberListByGroupId(groupID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userInfo, err := a.UserStorage.GetMultipleUserInfo(userIDs)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(userInfo)
	w.Write(res)
}
