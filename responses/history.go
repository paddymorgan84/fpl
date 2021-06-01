package responses

import "time"

// ManagerHistory is the response returned by the FPL API when the history endpoint is called
type ManagerHistory struct {
	Current []HistoricGameweek `json:"current"`
	Past    []PastSeason       `json:"past"`
	Chips   []Chip             `json:"chips"`
}

// HistoricGameweek contains the details for a specifc gameweek in the managers history
type HistoricGameweek struct {
	Event              int `json:"event"`
	Points             int `json:"points"`
	TotalPoints        int `json:"total_points"`
	Rank               int `json:"rank"`
	RankSort           int `json:"rank_sort"`
	OverallRank        int `json:"overall_rank"`
	Bank               int `json:"bank"`
	Value              int `json:"value"`
	EventTransfers     int `json:"event_transfers"`
	EventTransfersCost int `json:"event_transfers_cost"`
	PointsOnBench      int `json:"points_on_bench"`
}

// PastSeason contains the details for a past season in the managers history
type PastSeason struct {
	SeasonName  string `json:"season_name"`
	TotalPoints int    `json:"total_points"`
	Rank        int    `json:"rank"`
}

// Chip contains details around a chip and when it was used
type Chip struct {
	Name  string    `json:"name"`
	Time  time.Time `json:"time"`
	Event int       `json:"event"`
}
