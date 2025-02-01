package routes

import (
	"interviewPrep/restAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterConfig() *gin.Engine {
	router := gin.Default()
	router.GET("/events", getEvents)
	router.POST("/events", createEvent)
	return router
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse the event"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	event.Id = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create an event, try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!", "event": event})
}
