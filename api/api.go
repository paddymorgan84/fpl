package api

import (
	responses "github.com/paddymorgan84/fpl/responses"
)

// FplAPI is the interface for the FPL API
type FplAPI interface {
	GetBootstrapData() (responses.BootstrapData, error)
	GetAllFixtures() (responses.GameweekFixtures, error)
	GetGameweekFixtures(gameweek int) (responses.GameweekFixtures, error)
	GetGameweekPoints(teamID int, gameweek int) (responses.GameweekPoints, error)
	GetGameweekLiveScores(gameweek int) (responses.GameweekLiveScores, error)
	GetManagerHistory(teamID int) (responses.ManagerHistory, error)
	GetManagerDetails(teamID int) (responses.ManagerDetails, error)
	GetPlayerDetails(playerID int) (responses.PlayerDetails, error)
}

//go:generate mockgen -source=api.go -package=api -destination=mock_api.go
