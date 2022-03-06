package cmd

import (
	"fmt"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
)

// PointsArgs are the arguments you can pass to the points command
type PointsArgs struct {
	TeamID string
}

var pointsArgs PointsArgs

// BuildPointsCommand returns the points cobra command
func BuildPointsCommand(c api.FplAPI, config helpers.ConfigReader, teamsParser helpers.TeamsParser, gameweekParser helpers.GameweekParser, renderer ui.Renderer) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "points",
		Short: "Get the points for a specified gameweek (defaults to latest active gameweek)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getPoints(c, config, teamsParser, gameweekParser, renderer)
		},
	}

	var flags = cmd.Flags()
	flags.StringVarP(&pointsArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")

	return cmd
}

func getPoints(c api.FplAPI, config helpers.ConfigReader, teamsParser helpers.TeamsParser, gameweekParser helpers.GameweekParser, renderer ui.Renderer) error {
	teamID, err := teamsParser.GetTeamID(pointsArgs.TeamID, config)

	if err != nil {
		return err
	}

	bootstrap, err := c.GetBootstrapData()

	if err != nil {
		return err
	}

	gameweek, err := gameweekParser.GetCurrentGameweek(bootstrap.Gameweeks, config)

	if err != nil {
		return err
	}

	points, err := c.GetGameweekPoints(teamID, gameweek)

	if err != nil {
		return err
	}

	live, err := c.GetGameweekLiveScores(gameweek)

	if err != nil {
		return err
	}

	renderer.PrintHeader(fmt.Sprintf("Gameweek %d points", gameweek))
	renderer.PrintTeamPoints(bootstrap, live, points)

	return err
}
