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

// rivalsCmd represents the rivals command
var rivalsCmd = &cobra.Command{
	Use:   "rivals",
	Short: "Show the points for all of your rivals (specified in config) for a specified gameweek",
	Run: func(cmd *cobra.Command, args []string) {

		ui.PrintHeader("Rivals")

		if !viper.IsSet("rivals") {
			fmt.Println("No rivals specified. Update config for this to work.")
		}

		bootstrap := api.GetBootstrapData()
		gameweek := helpers.GetCurrentGameweek(bootstrap)
		live := api.GetLive(gameweek)

		for _, rival := range viper.GetStringSlice("rivals") {
			teamID, err := strconv.Atoi(rival)

			if err != nil {
				log.Fatal(err)
			}

			var points = api.GetPoints(teamID, gameweek)
			detailsResponse := api.GetDetails(teamID)
			ui.PrintRivalPoints(bootstrap, live, points, detailsResponse)
		}
	},
}

func init() {
	rootCmd.AddCommand(rivalsCmd)
}
