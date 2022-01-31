package controllers

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MaxScoreCalc(history *models.History) int {
	var maxScore int64
	database.DB.Model(&history).Select("max(score)").Where("user_id = ?", history.UserId).Find(&maxScore)
	return int(maxScore)
}

func AverageScoreCalc(history *models.History) float64 {
	var averageScore float64
	database.DB.Model(&history).Select("avg(score)").Where("user_id = ?", history.UserId).Find(&averageScore)
	return averageScore
}

func CountNumberOfTopCalc(history *models.History) int {
	var numberOfTopCount int64
	database.DB.Model(&history).Where(
		"rank = ? AND user_id = ?", 1, history.UserId).Count(&numberOfTopCount)
	return int(numberOfTopCount)
}

func CountNumberOfSecondsCalc(history *models.History) int {
	var numberOfSecondsCount int64
	database.DB.Model(&history).Where(
		"rank = ? AND user_id = ?", 2, history.UserId).Count(&numberOfSecondsCount)
	return int(numberOfSecondsCount)
}

func CountNumberOfThirdCalc(history *models.History) int {
	var numberOfThirdCount int64
	database.DB.Model(&history).Where(
		"rank = ? AND user_id = ?", 3, history.UserId).Count(&numberOfThirdCount)
	return int(numberOfThirdCount)
}

func CountNumberOfFourCalc(history *models.History) int {
	var numberOfFourCount int64
	database.DB.Model(&history).Where(
		"rank = ? AND user_id = ?", 4, history.UserId).Count(&numberOfFourCount)
	return int(numberOfFourCount)
}

func AverageRankCalc(history *models.History) float64 {
	var averageRank float64
	database.DB.Model(&history).Select("avg(rank)").Scan(&averageRank)
	return averageRank
}

func TopAverageRankCalc(history *models.History) float64 {
	var topAverageRank float64
	topAverageRank = float64(CountNumberOfTopCalc(history)) / float64(MatchCountCalc(history))
	return topAverageRank * 100.0
}

func FourthPlaceAvoidanceRateCalc(history *models.History) float64 {
	var fourthPlaceAvoidanceRate float64
	fourthPlaceAvoidanceRate = float64(CountNumberOfFourCalc(history)) / float64(MatchCountCalc(history))
	return 100.0 - fourthPlaceAvoidanceRate*100.0
}

func MatchCountCalc(history *models.History) int {
	var matchCount int64
	database.DB.Model(&history).Where("user_id", history.UserId).Count(&matchCount)
	return int(matchCount)
}

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

		var grade models.Grade

		grade = models.Grade{
			MaxScore:                 MaxScoreCalc(&history),
			AverageScore:             AverageScoreCalc(&history),
			NumberOfTop:              CountNumberOfTopCalc(&history),
			NumberOfSeconds:          CountNumberOfSecondsCalc(&history),
			NumberOfThird:            CountNumberOfThirdCalc(&history),
			NumberOfFour:             CountNumberOfFourCalc(&history),
			AverageRank:              AverageRankCalc(&history),
			TopAverageRank:           TopAverageRankCalc(&history),
			TotalPoints:              0.0,
			AveragePoints:            0.0,
			AdditionalTotalPoints:    0.0,
			AdditionalAveragePoints:  0.0,
			FourthPlaceAvoidanceRate: FourthPlaceAvoidanceRateCalc(&history),
			UserId:                   uint(Ids[i]),
			MatchCount:               MatchCountCalc(&history),
		}

		if database.DB.Model(&grade).Where("user_id = ?", history.UserId).Updates(&grade).RowsAffected == 0 {
			database.DB.Create(&grade)
		}
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
