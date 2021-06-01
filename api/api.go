package api

import (
	responses "github.com/paddymorgan84/fpl/responses"
)

// FplAPI is the interface for the FPL API
type FplAPI interface {
	GetBootstrapData() responses.BootstrapData
	GetAllFixtures() responses.GameweekFixtures
	GetGameweekFixtures(gameweek int) responses.GameweekFixtures
	GetGameweekPoints(teamID int, gameweek int) responses.GameweekPoints
	GetGameweekLiveScores(gameweek int) responses.GameweekLiveScores
	GetManagerHistory(teamID int) responses.ManagerHistory
	GetManagerDetails(teamID int) responses.ManagerDetails
	GetPlayerDetails(playerID int) responses.PlayerDetails
}
