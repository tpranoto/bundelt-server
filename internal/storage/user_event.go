package storage

type UserEventStorage interface {
	FindMemberListByEventID(eventID int64) (map[int64]UserEventRel, []int64, error)
	WriteUserEventRel(userID, eventID int64) error
}

type (
	UserEventRel struct {
		UserId  int64 `db:"user_id" json:"user_id"`
		EventId int64 `db:"event_id" json:"event_id"`
	}
)

func (p *postgreSQLStorage) FindMemberListByEventID(eventID int64) (map[int64]UserEventRel, []int64, error) {
	q := `SELECT event_id, user_id FROM event_user_rel WHERE event_id = $1`

	rows, err := p.db.QueryContext(p.ctx, q, eventID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	result := map[int64]UserEventRel{}
	userIDs := []int64{}
	for rows.Next() {
		var res UserEventRel

		err = rows.Scan(&res.EventId, &res.UserId)
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

func (p *postgreSQLStorage) WriteUserEventRel(userID, eventID int64) error {
	q := `INSERT INTO event_user_rel(user_id,event_id) VALUES($1,$2)`

	_, err := p.db.ExecContext(p.ctx, q, userID, eventID)
	if err != nil {
		return err
	}

	return nil
}
