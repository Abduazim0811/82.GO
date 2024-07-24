package api

import (
	"log"

	"82.GO/api/handler"
	"82.GO/internal/mongodb"
	"github.com/gin-gonic/gin"
)


func Routes(){
	router := gin.Default()
	db, err := mongodb.NewMessage()
	if err != nil {
		log.Fatal(err)
	}
	messagehandler := handler.NewHandler(db)

	router.POST("/message", messagehandler.CreateMessage)
	router.GET("/message/:id", messagehandler.GetbyIdMessage)
	router.Run(":8888")
}