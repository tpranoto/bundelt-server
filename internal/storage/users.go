package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tpranoto/bundelt-server/common"
)

type UsersStorage interface {
	CreateUserInfo(input CreateUserInfoInput) error
	UpdateUserInfo(input UpdateUserInfoInput) error
	GetUserInfo(userID int64) (UserInfo, error)
}

type (
	UserInfo struct {
		UserID          int64      `db:"user_id" json:"user_id"`
		Username        string     `db:"username" json:"username"`
		Email           string     `db:"email" json:"email"`
		MSISDN          string     `db:"msisdn" json:"msisdn"`
		Pwd             string     `db:"pwd" json:"pwd"`
		UserLocationStr []byte     `db:"user_loc" json:"-"`
		UserLocation    Location   `db:"-" json:"user_loc"`
		FullName        *string    `db:"full_name" json:"full_name"`
		Sex             *int       `db:"sex" json:"sex"`
		DOB             *time.Time `db:"dob" json:"dob"`
		LastLogin       time.Time  `db:"last_login" json:"last_login"`
		CreateTime      time.Time  `db:"create_time" json:"create_time"`
		UpdateTime      *time.Time `db:"update_time" json:"update_time"`
	}

	Location struct {
		Latitude  float64 `db:"latitude" json:"latitude"`
		Longitude float64 `db:"longitude" json:"longitude"`
	}

	CreateUserInfoInput struct {
		Username     string   `json:"username"`
		Email        string   `json:"email"`
		MSISDN       string   `json:"msisdn"`
		Pwd          string   `json:"pwd"`
		UserLocation Location `json:"user_loc"`
	}

	UpdateUserInfoInput struct {
		UserID          int64
		Username        common.UpdateString
		Email           common.UpdateString
		MSISDN          common.UpdateString
		Pwd             common.UpdateString
		UserLocationLat common.UpdateFloat64
		UserLocationLon common.UpdateFloat64
		FullName        common.UpdateString
		Sex             common.UpdateInt
	}
)

func (p postgreSQLStorage) GetUserInfo(userID int64) (userDetails UserInfo, err error) {
	query := fmt.Sprintf(`SELECT * FROM users WHERE user_id = %d`, userID)

	err = p.db.GetContext(p.ctx, &userDetails, query)
	if err != nil {
		return userDetails, err
	}

	loc := Location{}
	err = json.Unmarshal(userDetails.UserLocationStr, &loc)
	if err != nil {
		return userDetails, err
	}

	userDetails.UserLocation = loc
	return
}

func (p postgreSQLStorage) CreateUserInfo(userInfo CreateUserInfoInput) (err error) {
	q := `INSERT INTO users(username, email, pwd, msisdn, user_loc, last_login, create_time) 
	VALUES(?,?,?,?,?,NOW(),NOW()) RETURNING user_id`
	q = p.db.Rebind(q)

	locStr, err := json.Marshal(userInfo.UserLocation)
	if err != nil {
		return
	}

	var userID int64
	err = p.db.QueryRowContext(p.ctx, q, userInfo.Username, userInfo.Email, userInfo.MSISDN, userInfo.Pwd, locStr).Scan(&userID)
	if err != nil {
		return
	}
	p.logger.Println(userID)

	return
}

func (p postgreSQLStorage) UpdateUserInfo(input UpdateUserInfoInput) (err error) {
	return
}

func queryBuilder(query string, input interface{}) (string, []interface{}, error) {
	switch inp := input.(type) {
	case UpdateUserInfoInput:
		fields := ""
		values := ""
		params := []interface{}{}

		if inp.FullName.Update {
			fields += ", full_name"
			values += ",?"
			params = append(params, inp.FullName.Value)
		}

		if inp.Sex.Update {
			fields += ", sex"
			values += ",?"
			params = append(params, inp.Sex.Value)
		}

		return fmt.Sprintf(query, fields, values), params, nil
	}

	return "", nil, fmt.Errorf("error")
}
