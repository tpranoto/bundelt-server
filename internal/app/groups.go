package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tpranoto/bundelt-server/internal/storage"
)

func (a *App) handleGroupAdd(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.GroupInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	groupInfo, err := a.GroupStorage.WriteGroupDetails(input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(groupInfo)
	w.Write(res)
}

func (a *App) handleDeleteGroupDetail(w http.ResponseWriter, r *http.Request) {
	paramGroupID := r.FormValue("group_id")

	groupID, err := strconv.ParseInt(paramGroupID, 10, 64)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.GroupStorage.DeleteGroupDetails(groupID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handleNearbyGroups(w http.ResponseWriter, r *http.Request) {
	paramUser := r.FormValue("user_id")
	userID, _ := strconv.ParseInt(paramUser, 10, 64)
	paramLat := r.FormValue("lat")
	lat, _ := strconv.ParseFloat(paramLat, 64)
	paramLon := r.FormValue("lon")
	lon, _ := strconv.ParseFloat(paramLon, 64)
	paramLimit := r.FormValue("limit")
	limit, _ := strconv.Atoi(paramLimit)
	paramOffset := r.FormValue("offset")
	offset, _ := strconv.Atoi(paramOffset)

	groupRes, err := a.GroupStorage.GetNearbyGroups(lat, lon, limit, offset)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := []storage.GroupNearbyDetails{}
	for _, group := range groupRes {
		count, err := a.UserGroupRelStorage.FindMemberCountByGroupId(group.GroupID)
		if err != nil {
			a.Logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		joined, err := a.UserGroupRelStorage.CheckUserJoinedInGroup(group.GroupID, userID)
		if err != nil {
			a.Logger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		result = append(result, storage.GroupNearbyDetails{
			GroupID:   group.GroupID,
			Joined:    joined,
			GroupName: group.GroupName,
			Desc:      group.Desc,
			Created:   group.Created,
			Latitude:  group.Latitude,
			Longitude: group.Longitude,
			Distance:  group.Distance,
			Members:   count,
		})
	}

	res, _ := json.Marshal(result)
	w.Write(res)
}
