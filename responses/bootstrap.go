package responses

import "time"

// BootstrapResponse is the response returned when you request bootstrap data from the fpl API
type BootstrapResponse struct {
	Gameweeks    []Gameweek   `json:"events"`
	GameSettings GameSettings `json:"game_settings"`
	Phases       []Phase      `json:"phases"`
	Teams        []Team       `json:"teams"`
	TotalPlayers int          `json:"total_players"`
	Players      []Player     `json:"elements"`
	PlayerStats  []PlayerStat `json:"element_stats"`
	PlayerTypes  []PlayerType `json:"element_types"`
}

// Gameweek contains details around the gameweek
type Gameweek struct {
	ID                     int           `json:"id"`
	Name                   string        `json:"name"`
	DeadlineTime           time.Time     `json:"deadline_time"`
	AverageEntryScore      int           `json:"average_entry_score"`
	Finished               bool          `json:"finished"`
	DataChecked            bool          `json:"data_checked"`
	HighestScoringEntry    int           `json:"highest_scoring_entry"`
	DeadlineTimeEpoch      int           `json:"deadline_time_epoch"`
	DeadlineTimeGameOffset int           `json:"deadline_time_game_offset"`
	HighestScore           int           `json:"highest_score"`
	IsPrevious             bool          `json:"is_previous"`
	IsCurrent              bool          `json:"is_current"`
	IsNext                 bool          `json:"is_next"`
	ChipPlays              []ChipPlay    `json:"chip_plays"`
	MostSelected           int           `json:"most_selected"`
	MostTransferredIn      int           `json:"most_transferred_in"`
	TopPlayer              int           `json:"top_element"`
	TopPlayerInfo          TopPlayerInfo `json:"top_element_info"`
	TransfersMade          int           `json:"transfers_made"`
	MostCaptained          int           `json:"most_captained"`
	MostViceCaptained      int           `json:"most_vice_captained"`
}

// ChipPlay contains details around a chip that has been played
type ChipPlay struct {
	ChipName  string `json:"chip_name"`
	NumPlayed int    `json:"num_played"`
}

// TopPlayerInfo contains details around the top performing player
type TopPlayerInfo struct {
	ID     int `json:"id"`
	Points int `json:"points"`
}

// GameSettings contains details around the settings for the fpl game itself
type GameSettings struct {
	LeagueJoinPrivateMax         int           `json:"league_join_private_max"`
	LeagueJoinPublicMax          int           `json:"league_join_public_max"`
	LeagueMaxSizePublicClassic   int           `json:"league_max_size_public_classic"`
	LeagueMaxSizePublicH2H       int           `json:"league_max_size_public_h2h"`
	LeagueMaxSizePrivateH2H      int           `json:"league_max_size_private_h2h"`
	LeagueMaxKoRoundsPrivateH2H  int           `json:"league_max_ko_rounds_private_h2h"`
	LeaguePrefixPublic           string        `json:"league_prefix_public"`
	LeaguePointsH2HWin           int           `json:"league_points_h2h_win"`
	LeaguePointsH2HLose          int           `json:"league_points_h2h_lose"`
	LeaguePointsH2HDraw          int           `json:"league_points_h2h_draw"`
	LeagueKoFirstInsteadOfRandom bool          `json:"league_ko_first_instead_of_random"`
	CupStartEventID              int           `json:"cup_start_event_id"`
	CupStopEventID               int           `json:"cup_stop_event_id"`
	CupQualifyingMethod          string        `json:"cup_qualifying_method"`
	CupType                      string        `json:"cup_type"`
	SquadSquadplay               int           `json:"squad_squadplay"`
	SquadSquadsize               int           `json:"squad_squadsize"`
	SquadTeamLimit               int           `json:"squad_team_limit"`
	SquadTotalSpend              int           `json:"squad_total_spend"`
	UICurrencyMultiplier         int           `json:"ui_currency_multiplier"`
	UIUseSpecialShirts           bool          `json:"ui_use_special_shirts"`
	UISpecialShirtExclusions     []interface{} `json:"ui_special_shirt_exclusions"`
	StatsFormDays                int           `json:"stats_form_days"`
	SysViceCaptainEnabled        bool          `json:"sys_vice_captain_enabled"`
	TransfersCap                 int           `json:"transfers_cap"`
	TransfersSellOnFee           float64       `json:"transfers_sell_on_fee"`
	LeagueH2HTiebreakStats       []string      `json:"league_h2h_tiebreak_stats"`
	Timezone                     string        `json:"timezone"`
}

// Phase represents a phase in time
type Phase struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StartEvent int    `json:"start_event"`
	StopEvent  int    `json:"stop_event"`
}

// Team contains details around a specific fpl team
type Team struct {
	Code                int         `json:"code"`
	Draw                int         `json:"draw"`
	Form                interface{} `json:"form"`
	ID                  int         `json:"id"`
	Loss                int         `json:"loss"`
	Name                string      `json:"name"`
	Played              int         `json:"played"`
	Points              int         `json:"points"`
	Position            int         `json:"position"`
	ShortName           string      `json:"short_name"`
	Strength            int         `json:"strength"`
	TeamDivision        interface{} `json:"team_division"`
	Unavailable         bool        `json:"unavailable"`
	Win                 int         `json:"win"`
	StrengthOverallHome int         `json:"strength_overall_home"`
	StrengthOverallAway int         `json:"strength_overall_away"`
	StrengthAttackHome  int         `json:"strength_attack_home"`
	StrengthAttackAway  int         `json:"strength_attack_away"`
	StrengthDefenceHome int         `json:"strength_defence_home"`
	StrengthDefenceAway int         `json:"strength_defence_away"`
	PulseID             int         `json:"pulse_id"`
}

// Player contains details around a specific player
type Player struct {
	ChanceOfPlayingNextRound         int         `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound         int         `json:"chance_of_playing_this_round"`
	Code                             int         `json:"code"`
	CostChangeEvent                  int         `json:"cost_change_event"`
	CostChangeEventFall              int         `json:"cost_change_event_fall"`
	CostChangeStart                  int         `json:"cost_change_start"`
	CostChangeStartFall              int         `json:"cost_change_start_fall"`
	DreamteamCount                   int         `json:"dreamteam_count"`
	ElementType                      int         `json:"element_type"`
	EpNext                           string      `json:"ep_next"`
	EpThis                           string      `json:"ep_this"`
	EventPoints                      int         `json:"event_points"`
	FirstName                        string      `json:"first_name"`
	Form                             string      `json:"form"`
	ID                               int         `json:"id"`
	InDreamteam                      bool        `json:"in_dreamteam"`
	News                             string      `json:"news"`
	NewsAdded                        time.Time   `json:"news_added"`
	NowCost                          int         `json:"now_cost"`
	Photo                            string      `json:"photo"`
	PointsPerGame                    string      `json:"points_per_game"`
	SecondName                       string      `json:"second_name"`
	SelectedByPercent                string      `json:"selected_by_percent"`
	Special                          bool        `json:"special"`
	SquadNumber                      interface{} `json:"squad_number"`
	Status                           string      `json:"status"`
	Team                             int         `json:"team"`
	TeamCode                         int         `json:"team_code"`
	TotalPoints                      int         `json:"total_points"`
	TransfersIn                      int         `json:"transfers_in"`
	TransfersInEvent                 int         `json:"transfers_in_event"`
	TransfersOut                     int         `json:"transfers_out"`
	TransfersOutEvent                int         `json:"transfers_out_event"`
	ValueForm                        string      `json:"value_form"`
	ValueSeason                      string      `json:"value_season"`
	WebName                          string      `json:"web_name"`
	Minutes                          int         `json:"minutes"`
	GoalsScored                      int         `json:"goals_scored"`
	Assists                          int         `json:"assists"`
	CleanSheets                      int         `json:"clean_sheets"`
	GoalsConceded                    int         `json:"goals_conceded"`
	OwnGoals                         int         `json:"own_goals"`
	PenaltiesSaved                   int         `json:"penalties_saved"`
	PenaltiesMissed                  int         `json:"penalties_missed"`
	YellowCards                      int         `json:"yellow_cards"`
	RedCards                         int         `json:"red_cards"`
	Saves                            int         `json:"saves"`
	Bonus                            int         `json:"bonus"`
	Bps                              int         `json:"bps"`
	Influence                        string      `json:"influence"`
	Creativity                       string      `json:"creativity"`
	Threat                           string      `json:"threat"`
	IctIndex                         string      `json:"ict_index"`
	InfluenceRank                    int         `json:"influence_rank"`
	InfluenceRankType                int         `json:"influence_rank_type"`
	CreativityRank                   int         `json:"creativity_rank"`
	CreativityRankType               int         `json:"creativity_rank_type"`
	ThreatRank                       int         `json:"threat_rank"`
	ThreatRankType                   int         `json:"threat_rank_type"`
	IctIndexRank                     int         `json:"ict_index_rank"`
	IctIndexRankType                 int         `json:"ict_index_rank_type"`
	CornersAndIndirectFreekicksOrder interface{} `json:"corners_and_indirect_freekicks_order"`
	CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
	DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
	DirectFreekicksText              string      `json:"direct_freekicks_text"`
	PenaltiesOrder                   interface{} `json:"penalties_order"`
	PenaltiesText                    string      `json:"penalties_text"`
}

// PlayerStat contains a specific stat for a player
type PlayerStat struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

// PlayerType contains details around the player type
type PlayerType struct {
	ID                 int    `json:"id"`
	PluralName         string `json:"plural_name"`
	PluralNameShort    string `json:"plural_name_short"`
	SingularName       string `json:"singular_name"`
	SingularNameShort  string `json:"singular_name_short"`
	SquadSelect        int    `json:"squad_select"`
	SquadMinPlay       int    `json:"squad_min_play"`
	SquadMaxPlay       int    `json:"squad_max_play"`
	UIShirtSpecific    bool   `json:"ui_shirt_specific"`
	SubPositionsLocked []int  `json:"sub_positions_locked"`
	ElementCount       int    `json:"element_count"`
}
