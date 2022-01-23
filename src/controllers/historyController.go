package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AllHistories(c *gin.Context) {
	var histories []models.History

	database.DB.Preload("Match").Preload("User").Find(&histories)

	c.JSON(http.StatusOK, histories)

	return
}

func UserHistories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var histories []models.History

	database.DB.Where("user_id = ?", id).Preload("Match").Find(&histories)

	c.JSON(http.StatusOK, histories)

	return
}
