package models

import "time"

type Event struct {
	Id          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = make([]Event, 0)

func (e Event) Save() {
	// later add to the database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
