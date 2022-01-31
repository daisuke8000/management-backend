package models

type History struct {
	Model
	Score            int     `json:"score"`
	Rank             int     `json:"rank"`
	Points           float64 `json:"points"`
	AdditionalPoints float64 `json:"additional_points"`
	UserId           uint    `json:"user_id"`
	User             User    `json:"user" gorm:"foreignKey:UserId"`
	MatchId          uint    `json:"match_id"`
	Match            Match   `json:"match" gorm:"foreignKey:MatchId"`
}

const (
	BORDERSCORE  = 30000
	RANKFIRST    = 1
	RANKSECOND   = 2
	RANKTHIRD    = 3
	RANKFOURTH   = 4
	FIRSTBORNUS  = 50.0
	SECONDBORNUS = 10.0
	THIRDBORNUS  = -10.0
	FOURTHBORNUS = -30.0
)

func (h *History) CreateHistory(score, rank int, userId, matchId uint) {
	h.Score = score
	h.Rank = rank
	h.Points = (float64(score) - BORDERSCORE) / 1000.0
	h.AdditionalPoints = AdditionalPointsCalculate(
		rank, (float64(score)-BORDERSCORE)/1000.0)
	h.UserId = userId
	h.MatchId = matchId
	return
}

func AdditionalPointsCalculate(rank int, score float64) float64 {
	if rank == RANKFIRST {
		score += FIRSTBORNUS
	} else if rank == RANKSECOND {
		score += SECONDBORNUS
	} else if rank == RANKTHIRD {
		score += THIRDBORNUS
	} else if rank == RANKFOURTH {
		score += FOURTHBORNUS
	}
	return score
}
