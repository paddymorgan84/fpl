package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// BuildRivalsCommand returns the rivals cobra command
func BuildRivalsCommand(c *api.FplAPI) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rivals",
		Short: "Show the points for all of your rivals (specified in config) for a specified gameweek",
		Run: func(cmd *cobra.Command, args []string) {
			getRivals(*c)
		},
	}

	return cmd
}

func getRivals(c api.FplAPI) {
	ui.PrintHeader("Rivals")

	if !viper.IsSet("rivals") {
		fmt.Println("No rivals specified. Update config for this to work.")
	}

	bootstrap := c.GetBootstrapData()
	gameweek := helpers.GetCurrentGameweek(bootstrap)
	live := c.GetLive(gameweek)

	for _, rival := range viper.GetStringSlice("rivals") {
		teamID, err := strconv.Atoi(rival)

		if err != nil {
			log.Fatal(err)
		}

		var points = c.GetPoints(teamID, gameweek)
		detailsResponse := c.GetDetails(teamID)
		ui.PrintRivalPoints(bootstrap, live, points, detailsResponse)
	}
}
