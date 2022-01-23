package models

type Match struct {
	Model
	MatchUser01      int `json:"match_user_id_01"`
	MatchUser01Point int `json:"match_user_01_point"`
	MatchUser01Rank  int `json:"match_user_01_rank"`
	MatchUser02      int `json:"match_user_id_02"`
	MatchUser02Point int `json:"match_user_02_point"`
	MatchUser02Rank  int `json:"match_user_02_rank"`
	MatchUser03      int `json:"match_user_id_03"`
	MatchUser03Point int `json:"match_user_03_point"`
	MatchUser03Rank  int `json:"match_user_03_rank"`
	MatchUser04      int `json:"match_user_id_04"`
	MatchUser04Point int `json:"match_user_04_point"`
	MatchUser04Rank  int `json:"match_user_04_rank"`
}

func (m *Match) CreateHistory() (userIdsSlice [4]int, userPointsSlice [4]int, userRankSlice [4]int) {
	userIdsSlice = [4]int{m.MatchUser01, m.MatchUser02, m.MatchUser03, m.MatchUser04}
	userPointsSlice = [4]int{m.MatchUser01Point, m.MatchUser02Point, m.MatchUser03Point, m.MatchUser04Point}
	userRankSlice = [4]int{m.MatchUser01Rank, m.MatchUser02Rank, m.MatchUser03Rank, m.MatchUser04Rank}
	return userIdsSlice, userPointsSlice, userRankSlice
}
