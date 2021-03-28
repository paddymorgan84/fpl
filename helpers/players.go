package helpers

import "github.com/paddymorgan84/fpl/responses"

// DetermineCaptainFlag will work out whether or not the pick provided is the captain or vice, and if so return the appropriate string: (C) or (VC)
func DetermineCaptainFlag(pick responses.Pick) string {
	if pick.IsCaptain {
		return "(C)"
	}
	if pick.IsViceCaptain {
		return "(VC)"
	}

	return ""
}

// GetPlayerName will cross reference the pick with the bootstrap data to determine the players web appropriate name
func GetPlayerName(pick responses.Pick, bootstrap responses.BootstrapResponse) string {
	for _, player := range bootstrap.Players {
		if player.ID == pick.Element {
			return player.WebName
		}
	}

	return ""
}

// GetPoints will cross reference the pick with the live data to return the players points (with the captaincy multiplier factored in)
func GetPoints(pick responses.Pick, live responses.LiveResponse) int {
	for _, player := range live.Players {
		if pick.Element == player.ID {
			if pick.Multiplier == 0 {
				return player.Stats.TotalPoints
			}

			return player.Stats.TotalPoints * pick.Multiplier
		}
	}

	return 0
}
