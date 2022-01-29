package main

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"math/rand"
)

func main() {
	database.Connect()
	for i := 0; i < 100; i++ {
		match := models.Match{
			MatchUserId01:    rand.Intn(100) + 1,
			MatchUserScore01: rand.Intn(50000-40000) + 40000,
			MatchUserRank01:  1,
			MatchUserId02:    rand.Intn(100) + 1,
			MatchUserScore02: rand.Intn(40000-30000) + 30000,
			MatchUserRank02:  2,
			MatchUserId03:    rand.Intn(100) + 1,
			MatchUserScore03: rand.Intn(20000-10000) + 10000,
			MatchUserRank03:  3,
			MatchUserId04:    rand.Intn(100) + 1,
			MatchUserScore04: rand.Intn(15000-5000) + 5000,
			MatchUserRank04:  4,
		}

		database.DB.Create(&match)

		Ids, Scores, Ranks := match.CreateMatchParameters()

		for v := 0; v <= 3; {
			history := models.History{
				Score:  Scores[v],
				Rank:   Ranks[v],
				UserId: uint(Ids[v]),
				Match:  match,
			}
			database.DB.Create(&history)
			v++
		}
	}
}
