package storage

type UserGroupsStorage interface {
	WriteUserGroupRelation(input UserGroupRel) error
	FindGroupsByUserFbId(userFbId string) ([]string, error)
	DeleteUserGroup(input UserGroupRel) error
}

type (
	UserGroupRel struct {
		UserFbId string `db:"user_fb_id" json:"user_fb_id"`
		GroupId  string `db:"group_id" json:"group_id"`
	}

	UserGroupDetailInput struct {
		UserFbId  string `db:"user_fb_id" json:"user_fb_id"`
		GroupId   string `db:"group_id" json:"group_id"`
		GroupName string `db:"group_name" json:"group_name"`
		Desc      string `db:"group_desc" json:"desc"`
		Created   string `db:"created" json:"created"`
		Latitude  string `db:"lat" json:"lat"`
		Longitude string `db:"lon" json:"lon"`
	}
)

func (p postgreSQLStorage) WriteUserGroupRelation(input UserGroupRel) error {
	q := `INSERT INTO user_group_rel(user_fb_id, group_id) VALUES($1,$2)`

	_, err := p.db.ExecContext(p.ctx, q, input.UserFbId, input.GroupId)
	if err != nil {
		return err
	}

	return nil
}

func (p postgreSQLStorage) FindGroupsByUserFbId(userFbId string) ([]string, error) {
	q := `SELECT group_id FROM user_group_rel WHERE user_fb_id = $1`

	rows, err := p.db.QueryContext(p.ctx, q, userFbId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []string{}
	for rows.Next() {
		var res string

		err = rows.Scan(&res)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}

func (p postgreSQLStorage) DeleteUserGroup(input UserGroupRel) error {
	q := `DELETE FROM user_group_rel WHERE user_fb_id = $1 AND group_id = $2`

	_, err := p.db.ExecContext(p.ctx, q, input.UserFbId, input.GroupId)
	if err != nil {
		return err
	}

	return nil
}
