package models

type Grade struct {
	Model
	MaxScore                 int     `json:"max_score"`
	AverageScore             float64 `json:"average_score"`
	NumberOfTop              int     `json:"number_of_top"`
	NumberOfSeconds          int     `json:"number_of_seconds"`
	NumberOfThird            int     `json:"number_of_third"`
	NumberOfFour             int     `json:"number_of_four"`
	AverageRank              float64 `json:"average_rank"`
	TopAverageRank           float64 `json:"top_average_rank"`
	FourthPlaceAvoidanceRate float64 `json:"fourth_place_avoidance_rate"`
	UserId                   uint    `json:"user_id"`
	User                     User    `json:"user" gorm:"foreignKey:UserId"`
	MatchCount               int     `json:"match_count"`
}

//func (g *Grade) UpdateGrade(history *History) {
//	g.MaxScore = history.Score
//	g.AverageScore = 100
//	g.NumberOfTop = history.Rank
//	g.NumberOfSeconds = history.Rank
//	g.NumberOfThird = history.Rank
//	g.NumberOfFour = history.Rank
//	g.AverageRank = 0.1
//	g.TopAverageRank = 0.1
//	g.FourthPlaceAvoidanceRate = 0.1
//	g.MatchCount = 1
//	database.DB.Create(&g)
//	return
//}
