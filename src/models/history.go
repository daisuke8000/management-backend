package models

import "time"

type History struct {
	HistoryId uint      `json:"history_id"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Point     int       `json:"point"`
	Rank      int       `json:"rank"`
	//UserId       User         `json:"user_id" gorm:"foreignKey:UserId"`
	//MatchHistory MatchHistory `json:"match_history" gorm:"foreignKey:MatchHistoryId"`
}

type MatchHistory struct {
	MatchHistoryId uint      `json:"match_history_id"`
	Created        time.Time `json:"created"`
	Updated        time.Time `json:"updated"`
}
