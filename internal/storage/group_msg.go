package storage

import "time"

type GroupMessageStorage interface {
	WriteNewMessage(input GroupMsgInput) error
	GetMessageByGroupID(groupID int64) ([]GroupMsgInfo, []int64, error)
}

type (
	GroupMsgInput struct {
		GroupID int64  `db:"group_id" json:"group_id"`
		UserID  int64  `db:"user_id" json:"user_id"`
		Message string `db:"msg" json:"msg"`
	}

	GroupMsgInfo struct {
		GroupMsgID int64     `db:"group_msg_id" json:"group_msg_id"`
		GroupID    int64     `db:"group_id" json:"group_id"`
		UserID     int64     `db:"user_id" json:"user_id"`
		Message    string    `db:"msg" json:"msg"`
		CreateAt   time.Time `db:"create_time" json:"create_time"`
	}

	GroupMsgWithUserDetails struct {
		GroupMsgID int64  `db:"group_msg_id" json:"group_msg_id"`
		GroupID    int64  `db:"group_id" json:"group_id"`
		UserID     int64  `db:"user_id" json:"user_id"`
		FullName   string `db:"full_name" json:"full_name"`
		Message    string `db:"msg" json:"msg"`
		CreateAt   string `db:"create_time" json:"create_time"`
	}
)

func (p *postgreSQLStorage) WriteNewMessage(input GroupMsgInput) error {
	q := `INSERT INTO group_messages(group_id, user_id, msg, create_time) 
		VALUES($1,$2,$3,NOW())`

	_, err := p.db.ExecContext(p.ctx, q, input.GroupID, input.UserID, input.Message)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgreSQLStorage) GetMessageByGroupID(groupID int64) ([]GroupMsgInfo, []int64, error) {
	q := `SELECT * FROM group_messages WHERE group_id=$1 ORDER BY create_time DESC`

	rows, err := p.db.QueryContext(p.ctx, q, groupID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := []GroupMsgInfo{}
	userIDs := []int64{}
	for rows.Next() {
		var res GroupMsgInfo

		err = rows.Scan(&res.GroupMsgID, &res.GroupID, &res.UserID, &res.Message, &res.CreateAt)
		if err != nil {
			return nil, nil, err
		}

		result = append(result, res)
		userIDs = append(userIDs, res.UserID)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return result, userIDs, nil
}
