package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"82.GO/internal/models"
	"82.GO/internal/mongodb"
	"82.GO/internal/rabbitmq"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageHandler struct {
	db       *mongodb.MessageMongodb
	producer *rabbitmq.Producer
}

func NewHandler(db *mongodb.MessageMongodb, producer *rabbitmq.Producer) *MessageHandler {
	return &MessageHandler{db: db, producer: producer}
}

func (m *MessageHandler) CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.Timestamp = time.Now()

	body, err := json.Marshal(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := m.producer.PublishMessage("message_key", body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "message created")
}

func (m *MessageHandler) GetbyIdMessage(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	message, err := m.db.StoreGetbyIdMessage(objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
		return
	}

	c.JSON(http.StatusOK, message)
}
