package main

import (
	"github.com/daisuke8000/server/src/database"
	"github.com/daisuke8000/server/src/models"
	"math/rand"
)

func main() {
	database.Connect()
	for i := 0; i < 10; i++ {
		match := models.Match{
			MatchUser01:      rand.Intn(100-100) + 100,
			MatchUser01Point: rand.Intn(50000-40000) + 40000,
			MatchUser01Rank:  1,
			MatchUser02:      rand.Intn(100-100) + 100,
			MatchUser02Point: rand.Intn(40000-30000) + 30000,
			MatchUser02Rank:  2,
			MatchUser03:      rand.Intn(100-100) + 100,
			MatchUser03Point: rand.Intn(20000-10000) + 10000,
			MatchUser03Rank:  3,
			MatchUser04:      rand.Intn(100-100) + 100,
			MatchUser04Point: rand.Intn(15000-5000) + 5000,
			MatchUser04Rank:  4,
		}
		database.DB.Create(&match)
	}
}
