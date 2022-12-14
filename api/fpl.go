package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	responses "github.com/paddymorgan84/fpl/responses"
)

const endpoint = "https://fantasy.premierleague.com/api/"

// Client is the wrapper for the FplApi interface
type Client struct {
	Fpl *http.Client
}

// New will generate our FPL client to interact with the FPL API
func New() Client {
	client := Client{
		Fpl: &http.Client{},
	}

	return client
}

// GetBootstrapData returns all the bootstrap data that serves any additional calls
func (c Client) GetBootstrapData() (responses.BootstrapData, error) {
	var response responses.BootstrapData
	err := request("GET", endpoint+"bootstrap-static/", &response, *c.Fpl)

	if err != nil {
		return response, err
	}

	return response, err
}

// GetGameweekFixtures returns the fixtures for a specified gameweek
func (c Client) GetGameweekFixtures(gameweek int) (responses.GameweekFixtures, error) {
	var response responses.GameweekFixtures = make(responses.GameweekFixtures, 0)
	err := request("GET", endpoint+"fixtures/?event="+strconv.Itoa(gameweek), &response, *c.Fpl)

	return response, err
}

// GetAllFixtures returns the fixtures for every gameweek
func (c Client) GetAllFixtures() (responses.GameweekFixtures, error) {
	var response responses.GameweekFixtures = make(responses.GameweekFixtures, 0)
	err := request("GET", endpoint+"fixtures/", &response, *c.Fpl)

	return response, err
}

// GetGameweekPoints returns the points your team has for a specified gameweek
func (c Client) GetGameweekPoints(teamID int, gameweek int) (responses.GameweekPoints, error) {
	var response responses.GameweekPoints

	err := request("GET", endpoint+"entry/"+strconv.Itoa(teamID)+"/event/"+strconv.Itoa(gameweek)+"/picks/", &response, *c.Fpl)

	return response, err
}

// GetGameweekLiveScores returns player data for that specific gameweek
func (c Client) GetGameweekLiveScores(gameweek int) (responses.GameweekLiveScores, error) {
	var response responses.GameweekLiveScores

	err := request("GET", endpoint+"event/"+strconv.Itoa(gameweek)+"/live/", &response, *c.Fpl)

	return response, err
}

// GetManagerHistory returns the managers history for his team
func (c Client) GetManagerHistory(teamID int) (responses.ManagerHistory, error) {
	var response responses.ManagerHistory

	err := request("GET", endpoint+"entry/"+strconv.Itoa(teamID)+"/history/", &response, *c.Fpl)

	return response, err
}

// GetManagerDetails returns the details around the manager, more specifically this current season
func (c Client) GetManagerDetails(teamID int) (responses.ManagerDetails, error) {
	var response responses.ManagerDetails

	err := request("GET", endpoint+"entry/"+strconv.Itoa(teamID)+"/", &response, *c.Fpl)

	return response, err
}

// GetPlayerDetails returns the details for a specific player, including current and past seasons
func (c Client) GetPlayerDetails(playerID int) (responses.PlayerDetails, error) {
	var response responses.PlayerDetails

	err := request("GET", endpoint+"element-summary/"+strconv.Itoa(playerID)+"/", &response, *c.Fpl)

	return response, err
}

func request(method string, endpoint string, response interface{}, client http.Client) error {
	req, err := http.NewRequest(method, endpoint, nil)

	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", "")
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = resp.Body.Close()

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return err
	}

	return nil
}
