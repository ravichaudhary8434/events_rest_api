package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesaage": "Could not fetch event by id"})
		return
	}

	event, err := models.GetEventsByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not fetch event by id"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not register user for events"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId
	err := event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled"})
}
