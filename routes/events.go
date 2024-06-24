package routes

import (
	"net/http"
	"strconv"

	"github.com/Yadier01/golang-rest-api/models"
	"github.com/gin-gonic/gin"
)

func ShowEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not fetch event try again later"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func ShowEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "not working"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "cannot get event by id"})
		return
	}
	c.JSON(http.StatusOK, event)
}
func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "not working"})
		return
	}

	event.UserID = 1

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "could not save event", "error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "working"})

}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not save event", "error": err})
		return
	}
	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "not working"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "not working on update"})
		return
	}

  c.JSON(http.StatusOK, gin.H{"message": "Event Updated yurr"})
}
