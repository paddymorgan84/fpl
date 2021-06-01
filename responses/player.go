package responses

import "time"

// PlayerDetails returns details around a players upcoming fixtures, their gameweek history, and summary data for previous seasons
type PlayerDetails struct {
	Fixtures    []interface{}     `json:"fixtures"`
	History     []GameweekHistory `json:"history"`
	HistoryPast []PreviousSeasons `json:"history_past"`
}

// GameweekHistory contains the details for a player on a specific gameweek
type GameweekHistory struct {
	Element          int       `json:"element"`
	Fixture          int       `json:"fixture"`
	OpponentTeam     int       `json:"opponent_team"`
	TotalPoints      int       `json:"total_points"`
	WasHome          bool      `json:"was_home"`
	KickoffTime      time.Time `json:"kickoff_time"`
	TeamHScore       int       `json:"team_h_score"`
	TeamAScore       int       `json:"team_a_score"`
	Round            int       `json:"round"`
	Minutes          int       `json:"minutes"`
	GoalsScored      int       `json:"goals_scored"`
	Assists          int       `json:"assists"`
	CleanSheets      int       `json:"clean_sheets"`
	GoalsConceded    int       `json:"goals_conceded"`
	OwnGoals         int       `json:"own_goals"`
	PenaltiesSaved   int       `json:"penalties_saved"`
	PenaltiesMissed  int       `json:"penalties_missed"`
	YellowCards      int       `json:"yellow_cards"`
	RedCards         int       `json:"red_cards"`
	Saves            int       `json:"saves"`
	Bonus            int       `json:"bonus"`
	Bps              int       `json:"bps"`
	Influence        string    `json:"influence"`
	Creativity       string    `json:"creativity"`
	Threat           string    `json:"threat"`
	IctIndex         string    `json:"ict_index"`
	Value            int       `json:"value"`
	TransfersBalance int       `json:"transfers_balance"`
	Selected         int       `json:"selected"`
	TransfersIn      int       `json:"transfers_in"`
	TransfersOut     int       `json:"transfers_out"`
}

// PreviousSeasons contains summary details for a players previous seasons
type PreviousSeasons struct {
	SeasonName      string `json:"season_name"`
	ElementCode     int    `json:"element_code"`
	StartCost       int    `json:"start_cost"`
	EndCost         int    `json:"end_cost"`
	TotalPoints     int    `json:"total_points"`
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
}
