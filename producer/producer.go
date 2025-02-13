package producer

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
)

type RandomUserEvent struct {
	Name      string `json:"name"`
	Uid       string `json:"uid"`
	Timestamp string `json:"timeStamp"`
	EventType string `json:"eventType"`
	ProductId string `json:"productId"`
}

type ProductViewed struct {
	Name      string `json:"name"`
	Uid       string `json:"uid"`
	Timestamp string `json:"timeStamp"`
	EventType string `json:"eventType"`
	ProductId string `json:"productId"`
}

type ProductAddedToCart struct {
	Name      string `json:"name"`
	Uid       string `json:"uid"`
	Timestamp string `json:"timeStamp"`
	EventType string `json:"eventType"`
	ProductId string `json:"productId"`
}

func GenerateRandUserEvent() string {
	eventTypes := []string{"view", "added_to_cart", "removed_from_cart", "purchased", "shared", "searched"}
	randomEvent := rand.Intn(len(eventTypes))
	event := RandomUserEvent{
		Name:      faker.Name(),
		Uid:       faker.UUIDDigit(),
		Timestamp: time.Now().Format(time.RFC3339),
		EventType: eventTypes[randomEvent],
		ProductId: faker.UUIDDigit(),
	}

	jsonEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatal("error parsing json ", err)
	}

	return string(jsonEvent)
}
