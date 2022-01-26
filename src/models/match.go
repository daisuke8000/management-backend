package models

type Match struct {
	Model
	MatchUser01Id    int `json:"match_user01_id"`
	MatchUser01Point int `json:"match_user01_point"`
	MatchUser01Rank  int `json:"match_user01_rank"`
	MatchUser02Id    int `json:"match_user02_id"`
	MatchUser02Point int `json:"match_user02_point"`
	MatchUser02Rank  int `json:"match_user02_rank"`
	MatchUser03Id    int `json:"match_user03_id"`
	MatchUser03Point int `json:"match_user03_point"`
	MatchUser03Rank  int `json:"match_user03_rank"`
	MatchUser04Id    int `json:"match_user04_id"`
	MatchUser04Point int `json:"match_user04_point"`
	MatchUser04Rank  int `json:"match_user04_rank"`
}

func (m *Match) CreateHistory() (userIdsSlice [4]int, userPointsSlice [4]int, userRankSlice [4]int) {
	userIdsSlice = [4]int{m.MatchUser01Id, m.MatchUser02Id, m.MatchUser03Id, m.MatchUser04Id}
	userPointsSlice = [4]int{m.MatchUser01Point, m.MatchUser02Point, m.MatchUser03Point, m.MatchUser04Point}
	userRankSlice = [4]int{m.MatchUser01Rank, m.MatchUser02Rank, m.MatchUser03Rank, m.MatchUser04Rank}
	return userIdsSlice, userPointsSlice, userRankSlice
}
