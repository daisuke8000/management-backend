package models

type History struct {
	Model
	Score   int   `json:"score"`
	Rank    int   `json:"rank"`
	UserId  uint  `json:"user_id"`
	User    User  `json:"user" gorm:"foreignKey:UserId"`
	MatchId uint  `json:"match_id"`
	Match   Match `json:"match" gorm:"foreignKey:MatchId"`
}

func (h *History) CreateHistory(score, rank int, userId, matchId uint) {
	h.Score = score
	h.Rank = rank
	h.UserId = userId
	h.MatchId = matchId
	return
}
