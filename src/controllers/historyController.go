package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Histories(c *gin.Context) {
	var histories []models.History

	database.DB.Find(&histories)

	c.JSON(http.StatusOK, histories)

	return
}

func CreateHistory(c *gin.Context) {
	var history models.History

	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&history)

	c.JSON(http.StatusOK, history)

	return
}

func GetHistory(c *gin.Context) {
	var history models.History

	id, _ := strconv.Atoi(c.Param("id"))

	history.Id = uint(id)

	database.DB.Find(&history)

	c.JSON(http.StatusOK, history)

	return
}

func UpdateHistory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	history := models.History{}

	history.Id = uint(id)

	if err := c.ShouldBindJSON(&history); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&history).Updates(&history)

	c.JSON(http.StatusOK, history)

	return
}
