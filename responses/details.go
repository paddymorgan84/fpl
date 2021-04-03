package responses

import "time"

// DetailsResponse is the response returned by the FPL API when requesting from the manager details endpoint
type DetailsResponse struct {
	ID                         int       `json:"id"`
	JoinedTime                 time.Time `json:"joined_time"`
	StartedEvent               int       `json:"started_event"`
	FavouriteTeam              int       `json:"favourite_team"`
	PlayerFirstName            string    `json:"player_first_name"`
	PlayerLastName             string    `json:"player_last_name"`
	PlayerRegionID             int       `json:"player_region_id"`
	PlayerRegionName           string    `json:"player_region_name"`
	PlayerRegionIsoCodeShort   string    `json:"player_region_iso_code_short"`
	PlayerRegionIsoCodeLong    string    `json:"player_region_iso_code_long"`
	SummaryOverallPoints       int       `json:"summary_overall_points"`
	SummaryOverallRank         int       `json:"summary_overall_rank"`
	SummaryEventPoints         int       `json:"summary_event_points"`
	SummaryEventRank           int       `json:"summary_event_rank"`
	CurrentEvent               int       `json:"current_event"`
	Leagues                    League    `json:"leagues"`
	Name                       string    `json:"name"`
	Kit                        string    `json:"kit"`
	LastDeadlineBank           int       `json:"last_deadline_bank"`
	LastDeadlineValue          int       `json:"last_deadline_value"`
	LastDeadlineTotalTransfers int       `json:"last_deadline_total_transfers"`
}

// League holds the details for all the various league types available
type League struct {
	Classic []ClassicLeague `json:"classic"`
	H2H     []interface{}   `json:"h2h"`
	Cup     Cup             `json:"cup"`
}

// ClassicLeague holds all the details for a classic league
type ClassicLeague struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	ShortName      string      `json:"short_name"`
	Created        time.Time   `json:"created"`
	Closed         bool        `json:"closed"`
	Rank           interface{} `json:"rank"`
	MaxEntries     interface{} `json:"max_entries"`
	LeagueType     string      `json:"league_type"`
	Scoring        string      `json:"scoring"`
	AdminEntry     interface{} `json:"admin_entry"`
	StartEvent     int         `json:"start_event"`
	EntryCanLeave  bool        `json:"entry_can_leave"`
	EntryCanAdmin  bool        `json:"entry_can_admin"`
	EntryCanInvite bool        `json:"entry_can_invite"`
	HasCup         bool        `json:"has_cup"`
	CupLeague      interface{} `json:"cup_league"`
	CupQualified   interface{} `json:"cup_qualified"`
	EntryRank      int         `json:"entry_rank"`
	EntryLastRank  int         `json:"entry_last_rank"`
}

// Cup holds details on cup matches
type Cup struct {
	Matches   []Match `json:"matches"`
	Status    Status  `json:"status"`
	CupLeague int     `json:"cup_league"`
}

// Match holds specific match details for cup matches
type Match struct {
	ID               int         `json:"id"`
	Entry1Entry      int         `json:"entry_1_entry"`
	Entry1Name       string      `json:"entry_1_name"`
	Entry1PlayerName string      `json:"entry_1_player_name"`
	Entry1Points     int         `json:"entry_1_points"`
	Entry1Win        int         `json:"entry_1_win"`
	Entry1Draw       int         `json:"entry_1_draw"`
	Entry1Loss       int         `json:"entry_1_loss"`
	Entry1Total      int         `json:"entry_1_total"`
	Entry2Entry      int         `json:"entry_2_entry"`
	Entry2Name       string      `json:"entry_2_name"`
	Entry2PlayerName string      `json:"entry_2_player_name"`
	Entry2Points     int         `json:"entry_2_points"`
	Entry2Win        int         `json:"entry_2_win"`
	Entry2Draw       int         `json:"entry_2_draw"`
	Entry2Loss       int         `json:"entry_2_loss"`
	Entry2Total      int         `json:"entry_2_total"`
	IsKnockout       bool        `json:"is_knockout"`
	Winner           int         `json:"winner"`
	SeedValue        interface{} `json:"seed_value"`
	Event            int         `json:"event"`
	Tiebreak         interface{} `json:"tiebreak"`
}

// Status holds details around the status of the cup
type Status struct {
	QualificationEvent   int    `json:"qualification_event"`
	QualificationNumbers int    `json:"qualification_numbers"`
	QualificationRank    int    `json:"qualification_rank"`
	QualificationState   string `json:"qualification_state"`
}
