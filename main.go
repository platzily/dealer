package main

import (
	"context"
	"fmt"
	"time"

	"github.com/platzily/dealer/config"
	"github.com/platzily/dealer/domains"
	"github.com/platzily/dealer/drivers/mongodb"
)

func main() {

	mongoConfig := config.ReadMongoDBConfig()

	client := mongodb.NewConnection(mongoConfig.URL)
	defer client.Disconnect(context.TODO())

	eventRepository := mongodb.New(client)

	newEvent := domains.Event{
		Type:      "test",
		Payload:   "payloadTest",
		Sender:    "Some Service",
		State:     "Created",
		History:   []domains.EventStateHistory{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	eventRepository.Create(newEvent)

	result, _ := eventRepository.GetById(0)
	fmt.Printf("This is the result: %v", result)

	eventRepository.UpdateById(0, "Updated")

	updatedResult, _ := eventRepository.GetById(0)
	fmt.Printf("This is the updated result: %v", updatedResult)

	fmt.Println("Hello world, from notification dealer MF")

}
