package cmd

import (
	"fmt"

	fpl "github.com/paddymorgan84/fpl/api"
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
func BuildPointsCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "points",
		Short: "Get the points for a specified gameweek (defaults to latest active gameweek)",
		Run: func(cmd *cobra.Command, args []string) {
			teamID := helpers.GetTeamID(pointsArgs.TeamID)
			var bootstrap = fpl.GetBootstrapData()
			gameweek := helpers.GetCurrentGameweek(bootstrap)
			var points = fpl.GetPoints(teamID, gameweek)
			var live = fpl.GetLive(gameweek)

			ui.PrintHeader(fmt.Sprintf("Gameweek %d points", gameweek))
			ui.PrintTeamPoints(bootstrap, live, points)

		},
	}

	var flags = cmd.Flags()
	flags.StringVarP(&pointsArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")

	return cmd
}
