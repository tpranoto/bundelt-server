package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tpranoto/bundelt-server/internal/storage"
)

func (a *App) handlerAddEventGroupDetails(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.EventGroupDetail{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eventID, err := a.EventStorage.WriteEventDetails(storage.EventDetail{
		Name:      input.Name,
		Desc:      input.Desc,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	})
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.EventGroupStorage.WriteEventGroupRel(input.GroupID, eventID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.UserEventStorage.WriteUserEventRel(input.UserID, eventID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := input
	result.EventID = eventID

	res, _ := json.Marshal(result)
	w.Write(res)
}

func (a *App) handlerJoinEventGroup(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.EventGroup{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.EventGroupStorage.WriteEventGroupRel(input.GroupID, input.EventID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
