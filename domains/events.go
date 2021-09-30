package domains

type Event struct {
	Id      int64
	Type    string
	Payload interface{}
	Sender  string
	State   string
	history []Event
}

type EventModel interface {
	Create(e Event) error
	GetById(id int64) (Event, error)
}
