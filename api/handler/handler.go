package handler

import (
	"net/http"
	"time"

	"82.GO/internal/models"
	"82.GO/internal/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageHandler struct {
	db *mongodb.MessageMongodb
}

func NewHandler(db *mongodb.MessageMongodb) *MessageHandler {
	return &MessageHandler{db: db}
}

func (m *MessageHandler) CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.Timestamp = time.Now()

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
