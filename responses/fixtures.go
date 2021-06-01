package responses

import "time"

// GameweekFixtures is the response returned by the FPL API when requesting from the fixtures endpoint
type GameweekFixtures []struct {
	Code                 int       `json:"code"`
	Event                int       `json:"event"`
	Finished             bool      `json:"finished"`
	FinishedProvisional  bool      `json:"finished_provisional"`
	ID                   int       `json:"id"`
	KickoffTime          time.Time `json:"kickoff_time"`
	Minutes              int       `json:"minutes"`
	ProvisionalStartTime bool      `json:"provisional_start_time"`
	Started              bool      `json:"started"`
	AwayTeam             int       `json:"team_a"`
	AwayTeamScore        int       `json:"team_a_score"`
	HomeTeam             int       `json:"team_h"`
	HomeTeamScore        int       `json:"team_h_score"`
	Stats                []Stat    `json:"stats"`
	HomeTeamDifficulty   int       `json:"team_h_difficulty"`
	AwayTeamDifficulty   int       `json:"team_a_difficulty"`
	PulseID              int       `json:"pulse_id"`
}

// Stat holds a stat for the home and away team
type Stat struct {
	Identifier string        `json:"identifier"`
	A          []StatElement `json:"a"`
	H          []StatElement `json:"h"`
}

// StatElement holds a single stat value and element
type StatElement struct {
	Value   int `json:"value"`
	Element int `json:"element"`
}
