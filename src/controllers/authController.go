package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-contrib/sessions"
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

func Signin(c *gin.Context) {
	var data map[string]string
	session := sessions.Default(c)
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

	//payload := jwt.RegisteredClaims{
	//	Issuer: "develop",
	//	ExpiresAt: jwt.NewNumericDate(
	//		time.Unix(time.Now().Unix(), 0)),
	//}
	//
	//token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Invalid Credentials",
	//	})
	//	return
	//}
	//
	//c.SetSameSite(http.SameSiteNoneMode)
	//c.SetCookie(
	//	"jwt",
	//	token,
	//	3600,
	//	"/",
	//	"localhost",
	//	false,
	//	true,
	//)

	session.Set("loginUser", c.PostForm("userId"))
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}
