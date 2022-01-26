package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Matches(c *gin.Context) {

	var matches []models.Match

	id, _ := strconv.Atoi(c.Param("id"))

	database.DB.Find(&matches)

	database.DB.Where(
		"match_user01_id = ?", id).Or(
		"match_user02_id = ?", id).Or(
		"match_user03_id = ?", id).Or(
		"match_user04_id = ?", id).Find(&matches)

	c.JSON(http.StatusOK, matches)

	return
}

func CreateMatch(c *gin.Context) {
	var match models.Match

	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&match)

	c.JSON(http.StatusOK, match)

	Ids, Points, Ranks := match.CreateHistory()

	for i := 0; i <= 3; {
		history := models.History{
			Point:  Points[i],
			Rank:   Ranks[i],
			UserId: uint(Ids[i]),
			Match:  match,
		}
		database.DB.Create(&history)
		i++
	}

	return
}

func GetMatch(c *gin.Context) {
	var match models.Match

	id, _ := strconv.Atoi(c.Param("id"))

	match.Id = uint(id)

	database.DB.Find(&match)

	c.JSON(http.StatusOK, match)

	return
}

func UpdateMatch(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	match := models.Match{}

	match.Id = uint(id)

	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&match).Updates(&match)

	c.JSON(http.StatusOK, match)

	return
}
