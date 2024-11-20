package controllers

import (
	"net/http"
	"prayer-book/models"
	"prayer-book/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPrayers(c *gin.Context) {
	prayer, err := services.GetPrayers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch prayers"})
		return
	}

	c.JSON(http.StatusOK, prayer)
}

func GetPrayerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	prayers, err := services.GetPrayersByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	c.JSON(http.StatusOK, prayers)

}

func InsertPrayer(c *gin.Context) {
	var prayer models.PrayerStruct
	if err := c.ShouldBindJSON(&prayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.InsertPrayer(prayer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prayer successfully created",
		"id": id})
}
