package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/juju/ansiterm"
	fpl "github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// PointsArgs are the arguments you can pass to the points command
type PointsArgs struct {
	TeamID string
}

var pointsArgs PointsArgs

// pointsCmd represents the points command
var pointsCmd = &cobra.Command{
	Use:   "points",
	Short: "Get the points for a specified gameweek (defaults to latest active gameweek)",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		teamID := viper.GetInt("team-id")

		if teamID == 0 {
			teamID, err = strconv.Atoi(pointsArgs.TeamID)
			if err != nil {
				log.Fatal(err)
			}
		}
		var bootstrap = fpl.GetBootstrapData()
		var gameweek = 0

		gameweekParameter := viper.GetString("gameweek")

		if gameweekParameter == "" {
			gameweek = helpers.GetCurrentGameweek(bootstrap)
		} else {
			gameweek, err = strconv.Atoi(gameweekParameter)
			if err != nil {
				log.Fatal(err)
			}
		}

		var points = fpl.GetPoints(teamID, gameweek)
		var live = fpl.GetLive(gameweek)
		tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

		fmt.Printf("Gameweek %d points\n\n", gameweek)

		for _, pick := range points.Picks {
			captain := helpers.DetermineCaptainFlag(pick)
			name := helpers.GetPlayerName(pick, bootstrap)
			playerPoints := helpers.GetPoints(pick, live)

			fmt.Fprintf(tr, "%s %s\t%d\n", name, captain, playerPoints)
		}

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nTotal points: %d\n", points.EntryHistory.Points)
	},
}

func init() {
	rootCmd.AddCommand(pointsCmd)
	fixturesCmd.Flags().StringVarP(&pointsArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")
}
