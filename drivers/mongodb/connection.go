package mongodb

import (
	"context"
	"time"

	"github.com/platzily/dealer/config/mongodb"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection(uri string) (*mongo.Client, error) {

	url := mongodb.ReadMongoDBConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed disconnecting MongoDB connection: %s", err)
		}
	}()

	log.Infof("Mongo Connection Sucesss to %s", url)

	return client, nil
}
