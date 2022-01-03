package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Signup(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if data["password"] != data["password_confirm"] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password do not match",
		})
		return
	}
	user := models.User{
		Name:    data["name"],
		Email:   data["email"],
		IsAdmin: strings.Contains(c.FullPath(), "/api/admin"),
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)
	return
}
