package domains

import "time"

type Event struct {
	ID        int64               `bson:"_id"`
	Type      string              `bson:"type,omitempty"`
	Payload   interface{}         `bson:"payload,omitempty"`
	Sender    string              `bason:"sender,omitempty"`
	State     string              `bson:"state,omitempty"`
	History   []EventStateHistory `bson:"history"`
	CreatedAt time.Time           `bson:"created_at,omitempty"`
	UpdatedAt time.Time           `bson:"updated_at,omitempty"`
}

type EventStateHistory struct {
	State     string    `bson:"state,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
}

type EventModel interface {
	Create(e Event) error
	GetById(id int64) (Event, error)
	UpdateById(id int64, state string) error
}

type EventQueue interface {
	Subscribe(eventType string, callback func()) error
	Publish(event Event) error
}
