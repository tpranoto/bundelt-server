package storage

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tpranoto/bundelt-server/common"
)

type UsersStorage interface {
	CreateUserInfo(input UserInfoInput) (int64, error)
	UpdateUserInfo(input UpdateUserInfoInput) error
	GetUserInfo(userID int64) (UserInfoOutput, error)
	GetMultipleUserInfo(userIDs []int64) ([]UserInfoOutput, error)
	GetUserLogin(inp UserLoginInput) (UserInfoOutput, bool, error)
}

type (
	UserInfo struct {
		UserID     int64      `db:"user_id" json:"user_id"`
		Email      string     `db:"email" json:"email"`
		MSISDN     string     `db:"msisdn" json:"msisdn"`
		FullName   string     `db:"full_name" json:"full_name"`
		Pwd        string     `db:"pwd" json:"pwd"`
		Latitude   float64    `db:"lat" json:"lat"`
		Longitude  float64    `db:"lon" json:"lon"`
		LastLogin  time.Time  `db:"last_login" json:"last_login"`
		CreateTime time.Time  `db:"create_time" json:"create_time"`
		UpdateTime *time.Time `db:"update_time" json:"update_time"`
	}

	UserInfoInput struct {
		Email     string  `db:"email" json:"email"`
		MSISDN    string  `db:"msisdn" json:"msisdn"`
		FullName  string  `db:"full_name" json:"full_name"`
		Pwd       string  `db:"pwd" json:"pwd"`
		Latitude  float64 `db:"lat" json:"lat"`
		Longitude float64 `db:"lon" json:"lon"`
	}

	UserInfoOutput struct {
		UserID    int64   `db:"user_id" json:"user_id"`
		FullName  string  `db:"full_name" json:"full_name"`
		Latitude  float64 `db:"lat" json:"lat"`
		Longitude float64 `db:"lon" json:"lon"`
	}

	UserLoginInput struct {
		Email string  `json:"email"`
		Pwd   string  `json:"pwd"`
		Lat   float64 `json:"lat"`
		Lon   float64 `json:"lon"`
	}

	UpdateUserInfoInput struct {
		UserID    int64
		Email     common.UpdateString
		MSISDN    common.UpdateString
		Pwd       common.UpdateString
		Latitude  common.UpdateFloat64
		Longitude common.UpdateFloat64
		FullName  common.UpdateString
	}
)

func (p *postgreSQLStorage) GetUserInfo(userID int64) (userDetails UserInfoOutput, err error) {
	query := fmt.Sprintf(`SELECT * FROM users WHERE user_id = %d`, userID)

	err = p.db.GetContext(p.ctx, &userDetails, query)
	if err != nil {
		return userDetails, err
	}

	return
}

func (p *postgreSQLStorage) CreateUserInfo(userInfo UserInfoInput) (userID int64, err error) {
	q := `INSERT INTO users(email, msisdn, full_name, pwd, lat, lon, last_login, create_time) 
	VALUES($1,$2,$3,$4,$5,$6,NOW(),NOW()) RETURNING user_id`

	err = p.db.QueryRowContext(p.ctx, q, userInfo.Email, userInfo.MSISDN,
		userInfo.FullName, userInfo.Pwd, userInfo.Latitude, userInfo.Longitude).Scan(&userID)
	if err != nil {
		return
	}

	return
}

func (p *postgreSQLStorage) UpdateUserInfo(input UpdateUserInfoInput) (err error) {
	q := `UPDATE users SET %s WHERE user_id=?`

	query, args, err := updateQueryBuilder(q, input)
	if err != nil {
		return
	}

	query = p.db.Rebind(query)
	_, err = p.db.ExecContext(p.ctx, query, args...)
	if err != nil {
		return
	}

	return
}

func (p *postgreSQLStorage) GetUserLogin(inp UserLoginInput) (res UserInfoOutput, loginNotMatcing bool, err error) {
	q := `SELECT user_id, full_name, lat, lon FROM users WHERE email=$1 AND pwd=$2`

	err = p.db.QueryRowContext(p.ctx, q, inp.Email, inp.Pwd).Scan(&res.UserID, &res.FullName, &res.Latitude, &res.Longitude)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, true, nil
		}
		return
	}

	return
}

func (p *postgreSQLStorage) GetMultipleUserInfo(userIDs []int64) ([]UserInfoOutput, error) {
	if len(userIDs) <= 0 {
		return nil, nil
	}

	q := `SELECT user_id,full_name,lat,lon FROM users WHERE user_id IN (?)`

	query, args, err := sqlx.In(q, userIDs)
	if err != nil {
		return nil, err
	}
	query = p.db.Rebind(query)

	rows, err := p.db.QueryContext(p.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []UserInfoOutput{}
	for rows.Next() {
		var res UserInfoOutput

		err = rows.Scan(&res.UserID, &res.FullName, &res.Latitude, &res.Longitude)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func updateQueryBuilder(query string, input interface{}) (string, []interface{}, error) {
	switch inp := input.(type) {
	case UpdateUserInfoInput:
		fields := []string{}
		params := []interface{}{}

		if inp.Email.Update {
			fields = append(fields, "email=?")
			params = append(params, inp.Email.Value)
		}

		if inp.MSISDN.Update {
			fields = append(fields, "msisdn=?")
			params = append(params, inp.MSISDN.Value)
		}

		if inp.FullName.Update {
			fields = append(fields, "full_name=?")
			params = append(params, inp.FullName.Value)
		}

		if inp.Pwd.Update {
			fields = append(fields, "pwd=?")
			params = append(params, inp.Pwd.Value)
		}

		if inp.Latitude.Update {
			fields = append(fields, "lat=?")
			params = append(params, inp.Latitude.Value)
		}

		if inp.Longitude.Update {
			fields = append(fields, "lon=?")
			params = append(params, inp.Longitude.Value)
		}

		fields = append(fields, "update_time=NOW()")
		params = append(params, inp.UserID)

		return fmt.Sprintf(query, strings.Join(fields, ",")), params, nil
	}

	return "", nil, fmt.Errorf("error")
}
