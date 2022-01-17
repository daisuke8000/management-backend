package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/middleware"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func Signin(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong password",
		})
		return
	}

	payload := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	securedtype := false
	if host := c.Request.Host; string(host) == "localhost:8080" {
		//develop environment
		c.SetSameSite(http.SameSiteStrictMode)
	} else {
		//poduction environment
		c.SetSameSite(http.SameSiteNoneMode)
		securedtype = !securedtype
	}
	c.SetCookie(
		"jwt",
		token,
		3600,
		"/",
		"localhost",
		//develop environment >> secure: false
		//production environment >> secure: true
		securedtype,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}

func User(c *gin.Context) {

	id, _ := middleware.GetUserId(c)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	c.JSON(http.StatusOK, user)

	return
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"jwt",
		"",
		-3600,
		"/",
		"localhost",
		true,
		true,
	)
	c.JSON(http.StatusOK, gin.H{
		"message": "success Signout",
	})
	return
}

func UpdateInfo(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := middleware.GetUserId(c)
	user := models.User{
		Id:    id,
		Name:  data["name"],
		Email: data["email"],
	}

	database.DB.Model(&user).Updates(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

	return
}

func UpdatePassword(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if data["password"] != data["password_confirm"] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password do not match",
		})
		return
	}

	id, _ := middleware.GetUserId(c)

	user := models.User{
		Id: id,
	}

	database.DB.Where("id = ?", id).First(&user)

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

	return
}
