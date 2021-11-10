package storage

import "time"

type EventStorage interface {
	WriteEventDetails(input EventDetail) (int64, error)
}

type EventDetail struct {
	EventID   int64     `db:"event_id" json:"event_id"`
	Name      string    `db:"event_name" json:"event_name"`
	Desc      string    `db:"event_desc" json:"event_desc"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
	Latitude  float64   `db:"lat" json:"lat"`
	Longitude float64   `db:"lon" json:"lon"`
}

func (p *postgreSQLStorage) WriteEventDetails(input EventDetail) (int64, error) {
	q := `INSERT INTO events(event_name,event_desc,start_time,end_time,lat,lon) 
		VALUES($1,$2,$3,$4,$5,$6) RETURNING event_id`

	var eventID int64
	err := p.db.QueryRowContext(p.ctx, q, input.Name, input.Desc,
		input.StartTime, input.EndTime, input.Latitude, input.Longitude).Scan(&eventID)
	if err != nil {
		return eventID, err
	}

	return eventID, nil
}
