package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/juju/ansiterm"
	fpl "github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/responses"
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

		gameweekParamater := viper.GetString("gameweek")

		if gameweekParamater == "" {
			gameweek = getCurrentGameweek(bootstrap)
		} else {
			gameweek, err = strconv.Atoi(gameweekParamater)
			if err != nil {
				log.Fatal(err)
			}
		}

		var points = fpl.GetPoints(teamID, gameweek)
		var live = fpl.GetLive(gameweek)
		tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

		fmt.Printf("Gameweek %d points\n\n", gameweek)

		for _, pick := range points.Picks {
			captain := determineCaptainFlag(pick)
			name := getPlayerName(pick, bootstrap)
			playerPoints := getPoints(pick, live)

			fmt.Fprintf(tr, "%s %s\t%d\n", name, captain, playerPoints)
		}

		tr.Flush()

		fmt.Printf("\nTotal points: %d\n", points.EntryHistory.Points)
	},
}

func init() {
	rootCmd.AddCommand(pointsCmd)
	fixturesCmd.Flags().StringVarP(&pointsArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")
}

func determineCaptainFlag(pick responses.Pick) string {
	if pick.IsCaptain {
		return "(C)"
	}
	if pick.IsViceCaptain {
		return "(VC)"
	}

	return ""
}

func getPlayerName(pick responses.Pick, bootstrap responses.BootstrapResponse) string {
	for _, player := range bootstrap.Players {
		if player.ID == pick.Element {
			return player.WebName
		}
	}

	return ""
}

func getPoints(pick responses.Pick, live responses.LiveResponse) int {
	for _, player := range live.Players {
		if pick.Element == player.ID {
			return player.Stats.TotalPoints * pick.Multiplier
		}
	}

	return 0
}

func getCurrentGameweek(bootstrap responses.BootstrapResponse) int {
	for _, gameweek := range bootstrap.Gameweeks {
		if !gameweek.Finished {
			if gameweek.DeadlineTime.After(time.Now()) {
				return gameweek.ID - 1
			}

			return gameweek.ID
		}
	}

	return 38
}
