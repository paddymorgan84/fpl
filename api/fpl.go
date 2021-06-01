package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	responses "github.com/paddymorgan84/fpl/responses"
)

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
func (c Client) GetBootstrapData() responses.BootstrapData {
	var response responses.BootstrapData
	err := request("GET", "https://fantasy.premierleague.com/api/bootstrap-static/", &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetGameweekFixtures returns the fixtures for a specified gameweek
func (c Client) GetGameweekFixtures(gameweek int) responses.GameweekFixtures {
	var response responses.GameweekFixtures = make(responses.GameweekFixtures, 0)
	err := request("GET", "https://fantasy.premierleague.com/api/fixtures/?event="+strconv.Itoa(gameweek), &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetAllFixtures returns the fixtures for every gameweek
func (c Client) GetAllFixtures() responses.GameweekFixtures {
	var response responses.GameweekFixtures = make(responses.GameweekFixtures, 0)
	err := request("GET", "https://fantasy.premierleague.com/api/fixtures/", &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetGameweekPoints returns the points your team has for a specified gameweek
func (c Client) GetGameweekPoints(teamID int, gameweek int) responses.GameweekPoints {
	var response responses.GameweekPoints

	err := request("GET", "https://fantasy.premierleague.com/api/entry/"+strconv.Itoa(teamID)+"/event/"+strconv.Itoa(gameweek)+"/picks/", &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetGameweekLiveScores returns player data for that specific gameweek
func (c Client) GetGameweekLiveScores(gameweek int) responses.GameweekLiveScores {
	var response responses.GameweekLiveScores

	err := request("GET", "https://fantasy.premierleague.com/api/event/"+strconv.Itoa(gameweek)+"/live/", &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetManagerHistory returns the managers history for his team
func (c Client) GetManagerHistory(teamID int) responses.ManagerHistory {
	var response responses.ManagerHistory

	err := request("GET", "https://fantasy.premierleague.com/api/entry/"+strconv.Itoa(teamID)+"/history/", &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetManagerDetails returns the details around the manager, more specifically this current season
func (c Client) GetManagerDetails(teamID int) responses.ManagerDetails {
	var response responses.ManagerDetails

	err := request("GET", "https://fantasy.premierleague.com/api/entry/"+strconv.Itoa(teamID)+"/", &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

// GetPlayerDetails returns the details for a specific player, including current and past seasons
func (c Client) GetPlayerDetails(playerID int) responses.PlayerDetails {
	var response responses.PlayerDetails

	err := request("GET", "https://fantasy.premierleague.com/api/element-summary/"+strconv.Itoa(playerID)+"/", &response, *c.Fpl)

	if err != nil {
		log.Fatal(err)
	}

	return response
}

func request(method string, endpoint string, response interface{}, client http.Client) error {
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
