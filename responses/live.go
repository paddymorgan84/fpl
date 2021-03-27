package responses

type LiveResponse struct {
	Players []PlayerLive `json:"elements"`
}

type PlayerLive struct {
	ID      int           `json:"id"`
	Stats   Stats         `json:"stats"`
	Explain []ExplainStat `json:"explain"`
}

type Stats struct {
	Minutes         int    `json:"minutes"`
	GoalsScored     int    `json:"goals_scored"`
	Assists         int    `json:"assists"`
	CleanSheets     int    `json:"clean_sheets"`
	GoalsConceded   int    `json:"goals_conceded"`
	OwnGoals        int    `json:"own_goals"`
	PenaltiesSaved  int    `json:"penalties_saved"`
	PenaltiesMissed int    `json:"penalties_missed"`
	YellowCards     int    `json:"yellow_cards"`
	RedCards        int    `json:"red_cards"`
	Saves           int    `json:"saves"`
	Bonus           int    `json:"bonus"`
	Bps             int    `json:"bps"`
	Influence       string `json:"influence"`
	Creativity      string `json:"creativity"`
	Threat          string `json:"threat"`
	IctIndex        string `json:"ict_index"`
	TotalPoints     int    `json:"total_points"`
	InDreamteam     bool   `json:"in_dreamteam"`
}

type ExplainStat struct {
	Fixture int               `json:"fixture"`
	Stats   []ExplainStatStat `json:"stats"`
}

type ExplainStatStat struct {
	Identifier string `json:"identifier"`
	Points     int    `json:"points"`
	Value      int    `json:"value"`
}
