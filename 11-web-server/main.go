package main

import (
	"net/http"

	"github.com/Tyler-Liquornik/golang-fundamentals/11-web-server/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event

	// BindJSON auto-writes a 400 response and aborts on failure.
	// Alternatively, ShouldBindJSON returns an error for you to handle.
	// Missing fields are by default left as zero values in the struct instance.
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{"message:": "Event created", "event": event})
}
