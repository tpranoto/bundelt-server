package storage

import "database/sql"

type UserGroupsStorage interface {
	WriteUserGroupRelation(input UserGroupRel) error
	FindGroupsByUserId(userId int64) ([]int64, error)
	DeleteUserGroup(input UserGroupRel) error
	FindMemberCountByGroupId(groupId int64) (int64, error)
	FindMemberListByGroupId(groupId int64) (map[int64]UserGroupRel, []int64, error)
	CheckUserJoinedInGroup(groupId int64, userId int64) (bool, error)
}

type (
	UserGroupRel struct {
		UserId  int64 `db:"user_id" json:"user_id"`
		GroupId int64 `db:"group_id" json:"group_id"`
		Role    int   `db:"role" json:"role"`
	}

	UserGroupDetailInput struct {
		UserId    int64   `db:"user_id" json:"user_id"`
		GroupName string  `db:"group_name" json:"group_name"`
		Desc      string  `db:"group_desc" json:"desc"`
		Created   string  `db:"created" json:"created"`
		Latitude  float64 `db:"lat" json:"lat"`
		Longitude float64 `db:"lon" json:"lon"`
	}

	UserGroupDetailInfo struct {
		UserId    int64   `db:"user_id" json:"user_id"`
		GroupId   int64   `db:"group_id" json:"group_id"`
		GroupName string  `db:"group_name" json:"group_name"`
		Desc      string  `db:"group_desc" json:"desc"`
		Created   string  `db:"created" json:"created"`
		Latitude  float64 `db:"lat" json:"lat"`
		Longitude float64 `db:"lon" json:"lon"`
	}

	UserGroupRelWithGroupDetails struct {
		UserId    int64   `db:"user_id" json:"user_id"`
		GroupId   int64   `db:"group_id" json:"group_id"`
		Role      int     `db:"role" json:"role"`
		FullName  string  `db:"full_name" json:"full_name"`
		Latitude  float64 `db:"lat" json:"lat"`
		Longitude float64 `db:"lon" json:"lon"`
	}
)

func (p *postgreSQLStorage) WriteUserGroupRelation(input UserGroupRel) error {
	q := `INSERT INTO user_group_rel(user_id, group_id, role) VALUES($1,$2,$3)`

	_, err := p.db.ExecContext(p.ctx, q, input.UserId, input.GroupId, input.Role)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgreSQLStorage) FindGroupsByUserId(userId int64) ([]int64, error) {
	q := `SELECT group_id FROM user_group_rel WHERE user_id = $1`

	rows, err := p.db.QueryContext(p.ctx, q, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []int64{}
	for rows.Next() {
		var res int64

		err = rows.Scan(&res)
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

func (p *postgreSQLStorage) FindMemberCountByGroupId(groupId int64) (int64, error) {
	q := `SELECT COUNT(user_id) FROM user_group_rel WHERE group_id = $1`

	var result int64
	rows, err := p.db.QueryContext(p.ctx, q, groupId)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return result, err
		}
	}
	if err := rows.Err(); err != nil {
		return result, err
	}

	return result, nil
}

func (p *postgreSQLStorage) DeleteUserGroup(input UserGroupRel) error {
	q := `DELETE FROM user_group_rel WHERE user_id = $1 AND group_id = $2`

	_, err := p.db.ExecContext(p.ctx, q, input.UserId, input.GroupId)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgreSQLStorage) FindMemberListByGroupId(groupId int64) (map[int64]UserGroupRel, []int64, error) {
	q := `SELECT group_id, user_id, role FROM user_group_rel WHERE group_id = $1`

	rows, err := p.db.QueryContext(p.ctx, q, groupId)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := map[int64]UserGroupRel{}
	userIDs := []int64{}
	for rows.Next() {
		var res UserGroupRel

		err = rows.Scan(&res.GroupId, &res.UserId, &res.Role)
		if err != nil {
			return nil, nil, err
		}
		result[res.UserId] = res
		userIDs = append(userIDs, res.UserId)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return result, userIDs, nil
}

func (p postgreSQLStorage) CheckUserJoinedInGroup(groupId int64, userId int64) (res bool, err error) {
	q := `SELECT 1 FROM user_group_rel WHERE group_id = $1 AND user_id = $2`

	err = p.db.QueryRowContext(p.ctx, q, groupId, userId).Scan(&res)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	return res, nil
}
