package models

type Match struct {
	Model
	MatchUser01      uint `json:"match_user_id_01"`
	MatchUser01Point int  `json:"match_user_01_point"`
	MatchUser01Rank  int  `json:"match_user_01_rank"`
	MatchUser02      uint `json:"match_user_id_02"`
	MatchUser02Point int  `json:"match_user_02_point"`
	MatchUser02Rank  int  `json:"match_user_02_rank"`
	MatchUser03      uint `json:"match_user_id_03"`
	MatchUser03Point int  `json:"match_user_03_point"`
	MatchUser03Rank  int  `json:"match_user_03_rank"`
	MatchUser04      uint `json:"match_user_id_04"`
	MatchUser04Point int  `json:"match_user_04_point"`
	MatchUser04Rank  int  `json:"match_user_04_rank"`
}
