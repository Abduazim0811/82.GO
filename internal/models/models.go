package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Content    string             `json:"Content" bson:"content"`
	Timestamp  time.Time          `json:"timestamp" bson:"timestamp"`
	RoutingKey string             `json:"routingkey" bson:"routing_key"`
}
