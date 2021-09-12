package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tpranoto/bundelt-server/internal/storage"
)

func (a *App) handleAddNewGroupMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := storage.GroupMsgInput{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = a.GroupMessageStorage.WriteNewMessage(input)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) handleGetGroupMessage(w http.ResponseWriter, r *http.Request) {
	paramGroupID := r.FormValue("group_id")

	groupID, err := strconv.ParseInt(paramGroupID, 10, 64)
	if err != nil {
		a.Logger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, userIDs, err := a.GroupMessageStorage.GetMessageByGroupID(groupID)
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

	mappedUser := map[int64]storage.UserInfoOutput{}
	for _, user := range userInfo {
		mappedUser[user.UserID] = user
	}

	result := []storage.GroupMsgWithUserDetails{}
	for _, msg := range messages {
		result = append(result, storage.GroupMsgWithUserDetails{
			GroupMsgID: msg.GroupMsgID,
			GroupID:    msg.GroupID,
			Message:    msg.Message,
			UserID:     msg.UserID,
			FullName:   mappedUser[msg.UserID].FullName,
			CreateAt:   msg.CreateAt.Format("2006-01-02 15:04:05"),
		})
	}

	res, _ := json.Marshal(result)
	w.Write(res)
}
