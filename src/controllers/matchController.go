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
		"match_user_id01 = ?", id).Or(
		"match_user_id02 = ?", id).Or(
		"match_user_id03 = ?", id).Or(
		"match_user_id04 = ?", id).Find(&matches)

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

	Ids, Scores, Ranks := match.CreateMatchParameters()

	for i := 0; i <= 3; {
		var history models.History
		history.CreateHistory(Scores[i], Ranks[i], uint(Ids[i]), match.Id)
		database.DB.Create(&history)

		//var grade models.Grade
		//grade.UpdateGrade(&history)
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
