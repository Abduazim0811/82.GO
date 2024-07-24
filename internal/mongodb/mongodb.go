package mongodb

import (
	"context"
	"time"

	"82.GO/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMessage() (*MessageMongodb, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("test").Collection("messages")
	return &MessageMongodb{client: client, collection: collection}, nil
}

func (m *MessageMongodb) StoreNewMessage(message models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.collection.InsertOne(ctx, message)
	return err
}

func (m *MessageMongodb) StoreGetbyIdMessage(id primitive.ObjectID) (models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var message models.Message
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&message)
	return message, err
}
