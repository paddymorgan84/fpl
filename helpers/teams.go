package helpers

import (
	"strconv"

	"github.com/paddymorgan84/fpl/responses"
)

// GetTeamID will either convert the teamID provided to an int, or return the teamID stored in config under "team-id"
func GetTeamID(teamID string, config ConfigReader) (int, error) {
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
func GetTeam(teamID int, bootstrap responses.BootstrapData) string {
	for _, team := range bootstrap.Teams {
		if team.ID == teamID {
			return team.Name
		}
	}

	return ""
}

// CalculateMonetaryValue returns a correctly decimalised team value
func CalculateMonetaryValue(value int) float32 {
	return float32(value) / 10
}

// CalculateRankComparison returns the correct icon depending on rank change
func CalculateRankComparison(league responses.ClassicLeague) string {
	if league.EntryRank < league.EntryLastRank {
		return "🟢"
	}

	if league.EntryRank > league.EntryLastRank {
		return "🔴"
	}

	return "⚪"
}
