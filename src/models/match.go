package models

type Match struct {
	Model
	MatchUserId01    int `json:"match_user_id_01"`
	MatchUserScore01 int `json:"match_user_score_01"`
	MatchUserRank01  int `json:"match_user_rank_01"`
	MatchUserId02    int `json:"match_user_id_02"`
	MatchUserScore02 int `json:"match_user_score_02"`
	MatchUserRank02  int `json:"match_user_rank_02"`
	MatchUserId03    int `json:"match_user_id_03"`
	MatchUserScore03 int `json:"match_user_score_03"`
	MatchUserRank03  int `json:"match_user_rank_03"`
	MatchUserId04    int `json:"match_user_id_04"`
	MatchUserScore04 int `json:"match_user_score_04"`
	MatchUserRank04  int `json:"match_user_rank_04"`
}

func (m *Match) CreateMatchParameters() (userIdsSlice [4]int, userPointsSlice [4]int, userRankSlice [4]int) {
	userIdsSlice = [4]int{m.MatchUserId01, m.MatchUserId02, m.MatchUserId03, m.MatchUserId04}
	userPointsSlice = [4]int{m.MatchUserScore01, m.MatchUserScore02, m.MatchUserScore03, m.MatchUserScore04}
	userRankSlice = [4]int{m.MatchUserRank01, m.MatchUserRank02, m.MatchUserRank03, m.MatchUserRank04}
	return userIdsSlice, userPointsSlice, userRankSlice
}
