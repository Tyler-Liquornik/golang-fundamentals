package routes

import (
	"net/http"
	"strconv"

	"github.com/Tyler-Liquornik/golang-fundamentals/11-web-server/models"
	"github.com/gin-gonic/gin"
)

func createRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event ID"})
		return
	}

	// Make sure the event exists before trying to register a user for it.
	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = event.CreateRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create registration"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration created successfully"})
}

func deleteRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event ID"})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.DeleteRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete registration"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration deleted successfully"})
}
