package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (a *App) handlerGetEventMember(w http.ResponseWriter, r *http.Request) {
	paramEventId := r.FormValue("event_id")

	eventID, err := strconv.ParseInt(paramEventId, 10, 64)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, userIDs, err := a.UserEventStorage.FindMemberListByEventID(eventID)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(userIDs)
	w.Write(res)
}
