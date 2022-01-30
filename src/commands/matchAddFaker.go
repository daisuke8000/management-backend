package main

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"math/rand"
	"time"
)

const (
	RankOne    = 1
	RankTwo    = 2
	RankThree  = 3
	RankFour   = 4
	Percentage = 100.0
)

func main() {
	database.Connect()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		match := models.Match{
			MatchUserId01:    rand.Intn(300) + 23,
			MatchUserScore01: rand.Intn(50000-40000) + 40000,
			MatchUserRank01:  1,
			MatchUserId02:    rand.Intn(200) + 98,
			MatchUserScore02: rand.Intn(40000-30000) + 30000,
			MatchUserRank02:  2,
			MatchUserId03:    rand.Intn(200) + 11,
			MatchUserScore03: rand.Intn(20000-10000) + 10000,
			MatchUserRank03:  3,
			MatchUserId04:    rand.Intn(300) + 34,
			MatchUserScore04: rand.Intn(15000-5000) + 5000,
			MatchUserRank04:  4,
		}

		database.DB.Create(&match)

		Ids, Scores, Ranks := match.CreateMatchParameters()

		for v := 0; v <= 3; {
			var history models.History
			history.CreateHistory(Scores[v], Ranks[v], uint(Ids[v]), match.Id)
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
				FourthPlaceAvoidanceRate: FourthPlaceAvoidanceRateCalc(&history),
				UserId:                   uint(Ids[v]),
				MatchCount:               MatchCountCalc(&history),
			}

			if database.DB.Model(&grade).Where("user_id = ?", history.UserId).Updates(&grade).RowsAffected == 0 {
				database.DB.Create(&grade)
			}
			v++
		}
	}
}

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
		"rank = ? AND user_id = ?", RankOne, history.UserId).Count(&numberOfTopCount)
	return int(numberOfTopCount)
}

func CountNumberOfSecondsCalc(history *models.History) int {
	var numberOfSecondsCount int64
	database.DB.Model(&history).Where(
		"rank = ? AND user_id = ?", RankTwo, history.UserId).Count(&numberOfSecondsCount)
	return int(numberOfSecondsCount)
}

func CountNumberOfThirdCalc(history *models.History) int {
	var numberOfThirdCount int64
	database.DB.Model(&history).Where(
		"rank = ? AND user_id = ?", RankThree, history.UserId).Count(&numberOfThirdCount)
	return int(numberOfThirdCount)
}

func CountNumberOfFourCalc(history *models.History) int {
	var numberOfFourCount int64
	database.DB.Model(&history).Where(
		"rank = ? AND user_id = ?", RankFour, history.UserId).Count(&numberOfFourCount)
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
	return topAverageRank * Percentage
}

func FourthPlaceAvoidanceRateCalc(history *models.History) float64 {
	var fourthPlaceAvoidanceRate float64
	fourthPlaceAvoidanceRate = float64(CountNumberOfFourCalc(history)) / float64(MatchCountCalc(history))
	return Percentage - fourthPlaceAvoidanceRate*Percentage
}

func MatchCountCalc(history *models.History) int {
	var matchCount int64
	database.DB.Model(&history).Where("user_id", history.UserId).Count(&matchCount)
	return int(matchCount)
}
