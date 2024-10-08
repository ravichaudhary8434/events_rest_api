package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
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
	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesaage": "Could not parse the request", "error": err})
		return
	}

	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not create event", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesaage": "Could not fetch event by id"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventsByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not fetch event by id"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"mesaage": "Not authorized"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesaage": "Could not parse the request", "error": err})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not update event", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"mesaage": "Could not fetch event by id"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventsByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not fetch event by id"})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"mesaage": "Not authorized"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesaage": "Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
