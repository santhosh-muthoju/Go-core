package routes

import (
	"interviewPrep/restAPI/models"
	"log"
	"net/http"
	"time"

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
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not fetch the events!"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request data"})
		return
	}

	// Debug log to check received input
	log.Printf("Received Event: %+v\n", event)

	// Validate and Set Default Values
	if event.DateTime.IsZero() {
		event.DateTime = time.Now()
	}
	if event.UserId == 0 {
		event.UserId = 1
	}

	lastRtnId, err := models.Save(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not create the event!"})
		return
	}

	event.Id = lastRtnId
	log.Printf("Saved Event: %+v\n", event)

	context.JSON(http.StatusCreated, gin.H{
		"msg": "Event created!", "event": event})
}
