package responses

// GameweekLiveScores is the response returned by the FPL API to get the latest points for all players
type GameweekLiveScores struct {
	Players []PlayerLive `json:"elements"`
}

// PlayerLive contains the details of a player for that live gameweek
type PlayerLive struct {
	ID      int           `json:"id"`
	Stats   Stats         `json:"stats"`
	Explain []ExplainStat `json:"explain"`
}

// Stats represents a list of possible statistics for a player in the live gameweek
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

// ExplainStat gives a more detailed look at a specific stat for a specific fixture
type ExplainStat struct {
	Fixture int               `json:"fixture"`
	Stats   []ExplainStatStat `json:"stats"`
}

// ExplainStatStat holds details around the stat, including the value and how many points that attributes to
type ExplainStatStat struct {
	Identifier string `json:"identifier"`
	Points     int    `json:"points"`
	Value      int    `json:"value"`
}
