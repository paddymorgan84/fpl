package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/juju/ansiterm"
	"github.com/paddymorgan84/fpl/api"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// HistoryArgs are the arguments you can pass to the history command
type HistoryArgs struct {
	TeamID string
}

var historyArgs HistoryArgs

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Returns history for a managers current and past seasons",
	Run: func(cmd *cobra.Command, args []string) {
		teamID := 0
		var err error

		if historyArgs.TeamID == "" {
			teamID = viper.GetInt("team-id")
		} else {
			teamID, err = strconv.Atoi(historyArgs.TeamID)
			if err != nil {
				log.Fatal(err)
			}
		}

		p := message.NewPrinter(language.English)
		historyResponse := api.GetHistory(teamID)

		tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

		fmt.Println("------------")
		fmt.Println("This season")
		fmt.Println("------------")

		fmt.Fprintf(tr, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			"Gameweek",
			"Points",
			"Bench Points",
			"Gameweek Rank",
			"Transfers made",
			"Transfers cost",
			"Overall Points",
			"Overall Rank",
			"Team Value")

		for _, gameweek := range historyResponse.Current {
			_, err = p.Fprintf(tr, "%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%.1f\n",
				gameweek.Event,
				gameweek.Points,
				gameweek.PointsOnBench,
				gameweek.Rank,
				gameweek.EventTransfers,
				gameweek.EventTransfersCost,
				gameweek.TotalPoints,
				gameweek.OverallRank,
				helpers.CalculateMonetaryValue(gameweek.Value))

			if err != nil {
				log.Fatal(err)
			}
		}

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n------------")
		fmt.Println("Chips")
		fmt.Println("------------")

		fmt.Fprintf(tr, "%s\t%s\t%s\n", "Date", "Name", "Gamweeek used")

		for _, chip := range historyResponse.Chips {
			_, err = p.Fprintf(tr, "%s\t%s\t%d\n", chip.Time, chip.Name, chip.Event)

			if err != nil {
				log.Fatal(err)
			}
		}

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n------------")
		fmt.Println("Previous Seasons")
		fmt.Println("------------")

		fmt.Fprintf(tr, "%s\t%s\t%s\n", "Season", "Points", "Rank")

		for _, season := range historyResponse.Past {
			_, err = p.Fprintf(tr, "%s\t%d\t%d\n", season.SeasonName, season.TotalPoints, season.Rank)

			if err != nil {
				log.Fatal(err)
			}
		}

		err = tr.Flush()

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.Flags().StringVarP(&historyArgs.TeamID, "team-id", "t", "", "The team ID from FPL for your team")
}
