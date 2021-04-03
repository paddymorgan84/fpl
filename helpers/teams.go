package helpers

import "github.com/paddymorgan84/fpl/responses"

// GetTeam cross references the teamId with the bootstrap data to get the official team name
func GetTeam(teamID int, bootstrap responses.BootstrapResponse) string {
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
