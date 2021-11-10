package storage

import "time"

type EventGroupStorage interface {
	WriteEventGroupRel(groupID, eventID int64) error
}

type (
	EventGroupDetail struct {
		UserID    int64     `db:"user_id" json:"user_id"`
		GroupID   int64     `db:"group_id" json:"group_id"`
		EventID   int64     `db:"event_id" json:"event_id"`
		Name      string    `db:"event_name" json:"event_name"`
		Desc      string    `db:"event_desc" json:"event_desc"`
		Status    int       `db:"status" json:"status"`
		StartTime time.Time `db:"start_time" json:"start_time"`
		EndTime   time.Time `db:"end_time" json:"end_time"`
		Latitude  float64   `db:"lat" json:"lat"`
		Longitude float64   `db:"lon" json:"lon"`
	}

	EventGroup struct {
		GroupID int64 `db:"group_id" json:"group_id"`
		EventID int64 `db:"event_id" json:"event_id"`
	}
)

func (p *postgreSQLStorage) WriteEventGroupRel(groupID, eventID int64) error {
	q := `INSERT INTO event_group_rel(group_id,event_id) VALUES($1,$2)`

	_, err := p.db.ExecContext(p.ctx, q, groupID, eventID)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgreSQLStorage) GetEventMemberCount(eventID int64) (int64, error) {
	q := `SELECT COUNT(user_id) FROM event_group_rel WHERE event_id = $1`

	var result int64
	rows, err := p.db.QueryContext(p.ctx, q, eventID)
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
