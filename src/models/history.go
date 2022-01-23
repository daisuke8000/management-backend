package models

type History struct {
	Model
	Point   int   `json:"point"`
	Rank    int   `json:"rank"`
	UserId  uint  `json:"user_id"`
	User    User  `json:"user" gorm:"foreignKey:UserId"`
	MatchId uint  `json:"match_id"`
	Match   Match `json:"match" gorm:"foreignKey:MatchId"`
}
