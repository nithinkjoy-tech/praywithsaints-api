package controllers

import (
	"fmt"
	"net/http"
	"prayer-book/services"

	"github.com/gin-gonic/gin"
)

func GetPrayers(c *gin.Context) {
	fmt.Println("get prayer controller")

	prayer, err := services.GetPrayersService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch prayers"})
		return
	}

	c.JSON(http.StatusOK, prayer)
}

func GetPrayerById() {
	fmt.Println("get prayer by ID controller")
}

func CreatePrayer() {
	fmt.Println("create prayer controller")
}
