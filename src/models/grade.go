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
	TotalPoints              float64 `json:"total_points"`
	AveragePoints            float64 `json:"average_points"`
	AdditionalTotalPoints    float64 `json:"additional_total_points"`
	AdditionalAveragePoints  float64 `json:"additional_average_points"`
	FourthPlaceAvoidanceRate float64 `json:"fourth_place_avoidance_rate"`
	UserId                   uint    `json:"user_id"`
	User                     User    `json:"user" gorm:"foreignKey:UserId"`
	MatchCount               int     `json:"match_count"`
}
