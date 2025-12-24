package models

import (
	"rest-api-project/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {

	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	// Can also use Exec here but prepare prepares sometimes gives better performance
	// Also can be used with different values again
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	e.Id = id
	return err

}

func GetAllEvents() ([]Event, error) {

	var events = make([]Event, 0)
	query := "SELECT * FROM events"
	// Another option for executing the Query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
	 SELECT * FROM events WHERE id = ?
	`

	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (e Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)

	return err

}

func (e Event) DeleteEvent() error {
	query :=
		`
	DELETE FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id)

	return err

}

func (e Event) RegisterEvent(userId int64) error {

	query := `
	INSERT into registerations(user_id, event_id) 
	VALUES (?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id, userId)

	return err

}

func (e Event) CancelRegisteration(userId int64) error {
	query := `
	DELETE FROM registerations WHERE event_id = ? AND user_id = ?
	`
	_, err := db.DB.Query(query, e.Id, userId)

	return err

}
