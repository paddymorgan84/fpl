package api

import (
	responses "github.com/paddymorgan84/fpl/responses"
)

// FplAPI is the interface for the FPL API
type FplAPI interface {
	GetBootstrapData() responses.BootstrapResponse
	GetFixtures() responses.FixturesResponse
	GetPoints(teamID int, gameweek int) responses.PointsResponse
	GetLive(gameweek int) responses.LiveResponse
	GetHistory(teamID int) responses.HistoryResponse
	GetDetails(teamID int) responses.DetailsResponse
}
