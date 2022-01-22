package models

type History struct {
	Model
	TotalPoint int     `json:"total_point"`
	TotalRank  int     `json:"total_rank"`
	UserId     uint    `json:"user_id"`
	User       User    `json:"user" gorm:"foreignKey:UserId"`
	Matches    []Match `json:"matches,omitempty" gorm:"many2many:history_matches"`
}
