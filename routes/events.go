package routes

import (
	"example/event-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	// context.JSON(200, "Hello From REST API")
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in getting all Events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	searchId, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	event, err := models.GetEvent(searchId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event with the given ID not found"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing data"})
		return
	}

	event.UserId = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error while saving data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}

func updateEvent(context *gin.Context) {
	var newEvent models.Event
	err := context.ShouldBindBodyWithJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing data"})
		return
	}

	searchId, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	preEvent, err := models.GetEvent(searchId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event with the given ID not found"})
		return
	}

	if preEvent.UserId != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update this event"})
		return
	}
	newEvent.ID = preEvent.ID
	err = newEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error while updating data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated!"})
}

func deleteEvent(context *gin.Context) {
	searchId, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	event, err := models.GetEvent(searchId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event with the given ID not found"})
		return
	}

	if event.UserId != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error while deleting event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted!"})

}
