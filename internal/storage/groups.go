package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type GroupStorage interface {
	GetGroupDetailByIds(input []string) ([]GroupInfo, error)
	WriteGroupDetails(input GroupInfo) error
	DeleteGroupDetails(groupID string) error
	GetNearbyGroups(lat, lon float64, limit, offset int) ([]GroupInfoWithDistance, error)
}

type (
	GroupInfo struct {
		GroupID   string  `db:"group_id" json:"group_id"`
		GroupName string  `db:"group_name" json:"group_name"`
		Desc      string  `db:"group_desc" json:"desc"`
		Created   string  `db:"created" json:"created"`
		Latitude  float64 `db:"lat" json:"lat"`
		Longitude float64 `db:"lon" json:"lon"`
	}

	GroupInfoWithDistance struct {
		GroupID   string  `db:"group_id" json:"group_id"`
		GroupName string  `db:"group_name" json:"group_name"`
		Desc      string  `db:"group_desc" json:"desc"`
		Created   string  `db:"created" json:"created"`
		Latitude  float64 `db:"lat" json:"-"`
		Longitude float64 `db:"lon" json:"-"`
		Distance  float64 `db:"distance" json:"distance"`
	}
)

func (p postgreSQLStorage) GetGroupDetailByIds(groupIDs []string) ([]GroupInfo, error) {
	if len(groupIDs) <= 0 {
		return nil, nil
	}

	q := `SELECT * FROM groups WHERE group_id IN (?)`

	query, args, err := sqlx.In(q, groupIDs)
	if err != nil {
		return nil, err
	}
	query = p.db.Rebind(query)

	rows, err := p.db.QueryContext(p.ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []GroupInfo{}
	for rows.Next() {
		var res GroupInfo

		err = rows.Scan(&res.GroupID, &res.GroupName, &res.Desc, &res.Created, &res.Latitude, &res.Longitude)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}

func (p postgreSQLStorage) WriteGroupDetails(input GroupInfo) error {
	q := `INSERT INTO groups(group_id,group_name,group_desc,created,lat,lon) VALUES($1,$2,$3,$4,$5,$6)`

	_, err := p.db.ExecContext(p.ctx, q, input.GroupID, input.GroupName, input.Desc, input.Created, input.Latitude, input.Longitude)
	if err != nil {
		return err
	}
	return nil
}

func (p postgreSQLStorage) DeleteGroupDetails(groupID string) error {
	q := `DELETE FROM groups WHERE group_id = $1`

	_, err := p.db.ExecContext(p.ctx, q, groupID)
	if err != nil {
		return err
	}
	return nil
}

func (p postgreSQLStorage) GetNearbyGroups(lat, lon float64, limit, offset int) ([]GroupInfoWithDistance, error) {
	q := fmt.Sprintf(`SELECT group_id,group_name,group_desc,created,lat,lon,
			111.045* DEGREES(ACOS(LEAST(1.0, COS(RADIANS(latpoint))
                 * COS(RADIANS(lat))
                 * COS(RADIANS(longpoint) - RADIANS(lon))
                 + SIN(RADIANS(latpoint))
                 * SIN(RADIANS(lat))))) AS distance
			FROM groups
			JOIN(
				SELECT %f AS latpoint, %f AS longpoint
			) AS p ON 1=1
			ORDER BY distance
			LIMIT $1 OFFSET $2`, lat, lon)

	rows, err := p.db.QueryContext(p.ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []GroupInfoWithDistance{}
	for rows.Next() {
		var res GroupInfoWithDistance

		err = rows.Scan(&res.GroupID, &res.GroupName, &res.Desc, &res.Created, &res.Latitude, &res.Longitude, &res.Distance)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}
