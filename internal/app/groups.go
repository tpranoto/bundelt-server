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
		return
	}

	input := storage.GroupInfo{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	err = a.GroupStorage.WriteGroupDetails(input)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handleDeleteGroupDetail(w http.ResponseWriter, r *http.Request) {
	paramGroupID := r.FormValue("group_id")

	err := a.GroupStorage.DeleteGroupDetails(paramGroupID)
	if err != nil {
		a.Logger.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (a *App) handleNearbyGroups(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	res, _ := json.Marshal(groupRes)
	w.Write(res)
	w.WriteHeader(http.StatusOK)
}
