package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ambassadors(c *gin.Context) {
	var users []models.User

	database.DB.Where("is_admin = false").Find(&users)

	c.JSON(http.StatusOK, users)

	return
}
