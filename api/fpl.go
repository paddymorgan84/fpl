package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	responses "github.com/paddymorgan84/fpl/responses"
)

// GetBootstrapData returns all the bootstrap data that serves any additional calls
func GetBootstrapData() responses.BootstrapResponse {
	var response responses.BootstrapResponse
	err := request("GET", "https://fantasy.premierleague.com/api/bootstrap-static/", &response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetFixtures returns the fixtures for a specified gameweek
func GetFixtures() responses.FixturesResponse {
	var response responses.FixturesResponse = make(responses.FixturesResponse, 0)
	err := request("GET", "https://fantasy.premierleague.com/api/fixtures/", &response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetPoints returns the points your team has for a specified gameweek
func GetPoints(teamID int, gameweek int) responses.PointsResponse {
	var response responses.PointsResponse

	err := request("GET", "https://fantasy.premierleague.com/api/entry/"+strconv.Itoa(teamID)+"/event/"+strconv.Itoa(gameweek)+"/picks/", &response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetLive returns player data for that specific gameweek
func GetLive(gameweek int) responses.LiveResponse {
	var response responses.LiveResponse

	err := request("GET", "https://fantasy.premierleague.com/api/event/"+strconv.Itoa(gameweek)+"/live/", &response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetHistory returns the managers history for his team
func GetHistory(teamID int) responses.HistoryResponse {
	var response responses.HistoryResponse

	err := request("GET", "https://fantasy.premierleague.com/api/entry/"+strconv.Itoa(teamID)+"/history/", &response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetDetails returns the details around the manager, more specifically this current season
func GetDetails(teamID int) responses.DetailsResponse {
	var response responses.DetailsResponse

	err := request("GET", "https://fantasy.premierleague.com/api/entry/"+strconv.Itoa(teamID)+"/", &response)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func request(method string, endpoint string, response interface{}) error {
	var client = &http.Client{}
	req, err := http.NewRequest(method, endpoint, nil)

	req.Header.Add("User-Agent", "")
	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return err
	}

	return nil
}
