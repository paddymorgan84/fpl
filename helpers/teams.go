package helpers

import (
	"strconv"

	"github.com/paddymorgan84/fpl/responses"
)

// TeamsParser helps me abstract away the team related work so I can better unit test it
type TeamsParser interface {
	GetTeamID(teamID string, config ConfigReader) (int, error)
	GetTeam(teamID int, bootstrap responses.BootstrapData) string
	CalculateMonetaryValue(value int) float32
	CalculateRankComparison(league responses.ClassicLeague) string
}

// FplTeamsParser handles any team related parsing that needs to be done
type FplTeamsParser struct {
}

// GetTeamID will either convert the teamID provided to an int, or return the teamID stored in config under "team-id"
func (t FplTeamsParser) GetTeamID(teamID string, config ConfigReader) (int, error) {
	if teamID == "" {
		return config.GetInt("team-id"), nil
	}

	team, err := strconv.Atoi(teamID)

	if err != nil {
		return 0, err
	}

	return team, nil
}

// GetTeam cross references the teamId with the bootstrap data to get the official team name
func (t FplTeamsParser) GetTeam(teamID int, bootstrap responses.BootstrapData) string {
	for _, team := range bootstrap.Teams {
		if team.ID == teamID {
			return team.Name
		}
	}

	return ""
}

// CalculateMonetaryValue returns a correctly decimalised team value
func (t FplTeamsParser) CalculateMonetaryValue(value int) float32 {
	return float32(value) / 10
}

// CalculateRankComparison returns the correct icon depending on rank change
func (t FplTeamsParser) CalculateRankComparison(league responses.ClassicLeague) string {
	if league.EntryRank < league.EntryLastRank {
		return "🟢"
	}

	if league.EntryRank > league.EntryLastRank {
		return "🔴"
	}

	return "⚪"
}

//go:generate mockgen -source=teams.go -package=helpers -destination=mock_teams_parser.go
