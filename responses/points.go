package responses

// PointsResponse is the response returned by the FPL APi when the picks endpoint is called
type PointsResponse struct {
	ActiveChip    interface{}    `json:"active_chip"`
	AutomaticSubs []AutomaticSub `json:"automatic_subs"`
	EntryHistory  EntryHistory   `json:"entry_history"`
	Picks         []Pick         `json:"picks"`
}

// AutomaticSub identifies any subs that were automatically made if a first team player didn't play
type AutomaticSub struct {
	Entry      int `json:"entry"`
	ElementIn  int `json:"element_in"`
	ElementOut int `json:"element_out"`
	Event      int `json:"event"`
}

// EntryHistory gives details around that gameweeks picks
type EntryHistory struct {
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

// Pick contains information around a specific picked player for that gameweek
type Pick struct {
	Element       int  `json:"element"`
	Position      int  `json:"position"`
	Multiplier    int  `json:"multiplier"`
	IsCaptain     bool `json:"is_captain"`
	IsViceCaptain bool `json:"is_vice_captain"`
}
