package domains

type Event struct {
	ID      int64       `bson:"_id,omitempty"`
	Type    string      `bson:"type,omitempty"`
	Payload interface{} `bson:"payload,omitempty"`
	Sender  string      `bason:"sender,omitempty"`
	State   string      `bson:"state,omitempty"`
	history []Event     `bson:"history"`
}

type EventModel interface {
	Create(e Event) error
	GetById(id int64) (Event, error)
}
