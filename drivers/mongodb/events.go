package mongodb

import (
	"context"
	"time"

	"github.com/platzily/dealer/config"
	"github.com/platzily/dealer/domains"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const EVENTS_COLLECTION = "events"

var env *config.MongoDBConfig = config.ReadMongoDBConfig()

type EventRepository struct {
	conn *mongo.Client
}

func New(db *mongo.Client) domains.EventModel {

	return &EventRepository{
		conn: db,
	}
}

func (er *EventRepository) Create(e domains.Event) error {

	ctxOperation, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := er.conn.Database(env.Database).Collection(EVENTS_COLLECTION)

	res, err := collection.InsertOne(ctxOperation, e)

	if err != nil {
		log.Errorf("Error inserting event: %v", err)
		return err
	}

	log.Infof("Inserted a document: %v", res.InsertedID)

	return nil
}

func (er *EventRepository) GetById(id int64) (domains.Event, error) {

	ctxOperation, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := er.conn.Database(env.Database).Collection(EVENTS_COLLECTION)

	var result domains.Event
	err := collection.FindOne(ctxOperation, bson.D{{"_id", id}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Errorf("No documents found with the given id: %v", err)
			return domains.Event{}, err
		}
		log.Errorf("Error finding event: %v", err)
		return domains.Event{}, err
	}

	return result, nil
}
