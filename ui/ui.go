package ui

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/juju/ansiterm"
	"github.com/paddymorgan84/fpl/helpers"
	"github.com/paddymorgan84/fpl/responses"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Renderer is an abstraction that allows me to implement multiple renderers whilst also allowing me to unit test
type Renderer interface {
	PrintHeader(title string)
	PrintTeamPoints(bootstrap responses.BootstrapData, live responses.GameweekLiveScores, points responses.GameweekPoints)
	PrintRivalPoints(bootstrap responses.BootstrapData, live responses.GameweekLiveScores, points responses.GameweekPoints, details responses.ManagerDetails)
	PrintManagerDetails(details responses.ManagerDetails)
	PrintClassicLeagues(details responses.ManagerDetails, teamsParser helpers.TeamsParser)
	PrintGlobalLeagues(detailsResponse responses.ManagerDetails, teamsParser helpers.TeamsParser)
	PrintTransfersAndFinance(detailsResponse responses.ManagerDetails, teamsParser helpers.TeamsParser)
	PrintGameweekFixtures(bootstrap responses.BootstrapData, fixtures responses.GameweekFixtures, teamsParser helpers.TeamsParser, gameweek int)
	PrintSeasonDetails(history responses.ManagerHistory, teamsParser helpers.TeamsParser)
	PrintChipDetails(history responses.ManagerHistory)
	PrintPreviousSeasonDetails(history responses.ManagerHistory)
	PrintNoRivalsWarning()
}

// TerminalRenderer is my concrete implementation of the Renderer interface, outputting to the terminal
type TerminalRenderer struct {
}

// PrintHeader will print standardised header, calculating the appropriate length from the titles characters
func (t TerminalRenderer) PrintHeader(title string) {

	var b bytes.Buffer

	for i := 0; i <= len(title); i++ {
		b.WriteString("-")
	}

	fmt.Printf("\n%s\n", b.String())
	fmt.Printf("%s\n", title)
	fmt.Printf("%s\n\n", b.String())
}

// PrintTeamPoints prints the points the team has
func (t TerminalRenderer) PrintTeamPoints(bootstrap responses.BootstrapData, live responses.GameweekLiveScores, points responses.GameweekPoints) {
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	for _, pick := range points.Picks {
		captain := helpers.DetermineCaptainFlag(pick)
		name := helpers.GetPlayerName(pick, bootstrap)
		playerPoints := helpers.GetPoints(pick, live)

		fmt.Fprintf(tr, "%s %s\t%d\n", name, captain, playerPoints)
	}

	helpers.AutoFlush(tr)

	fmt.Printf("\nTotal points: %d\n", points.EntryHistory.Points)
}

// PrintRivalPoints prints out the details for the rivals given the teamID in details
func (t TerminalRenderer) PrintRivalPoints(bootstrap responses.BootstrapData, live responses.GameweekLiveScores, points responses.GameweekPoints, details responses.ManagerDetails) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	_, err := p.Fprintf(tr, "%s\t%s\n", "Team Name: ", details.Name)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t%s %s\n", "Manager: ", details.PlayerFirstName, details.PlayerLastName)

	if err != nil {
		log.Fatal(err)
	}

	for _, pick := range points.Picks {
		captain := helpers.DetermineCaptainFlag(pick)
		name := helpers.GetPlayerName(pick, bootstrap)
		playerPoints := helpers.GetPoints(pick, live)

		fmt.Fprintf(tr, "%s %s\t%d\n", name, captain, playerPoints)
	}

	helpers.AutoFlush(tr)

	fmt.Printf("\nTotal points: %d\n\n\n", points.EntryHistory.Points)
}

// PrintManagerDetails prints the summary details for a manager
func (t TerminalRenderer) PrintManagerDetails(details responses.ManagerDetails) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	_, err := p.Fprintf(tr, "%s\t%s\n", "Team Name: ", details.Name)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t%d\n", "Overall Points: ", details.SummaryOverallPoints)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t%d\n", "Overall Rank: ", details.SummaryOverallRank)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t%d\n", "Gameweek Points: ", details.SummaryEventPoints)

	if err != nil {
		log.Fatal(err)
	}

	helpers.AutoFlush(tr)
}

// PrintClassicLeagues prints all classic leagues and the current rank for each
func (t TerminalRenderer) PrintClassicLeagues(details responses.ManagerDetails, teamsParser helpers.TeamsParser) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	for _, league := range details.Leagues.Classic {
		if league.LeagueType == "x" {
			_, err := p.Fprintf(tr, "%s\t%d\t%s\n", league.Name, league.EntryRank, teamsParser.CalculateRankComparison(league))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	helpers.AutoFlush(tr)
}

// PrintGlobalLeagues prints all global leagues and the current rank for each
func (t TerminalRenderer) PrintGlobalLeagues(detailsResponse responses.ManagerDetails, teamsParser helpers.TeamsParser) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	for _, league := range detailsResponse.Leagues.Classic {
		if league.LeagueType == "s" {
			_, err := p.Fprintf(tr, "%s\t%d\t%s\n", league.Name, league.EntryRank, teamsParser.CalculateRankComparison(league))

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	helpers.AutoFlush(tr)
}

// PrintTransfersAndFinance prints details for a teamsd transfers, value and money in the bank
func (t TerminalRenderer) PrintTransfersAndFinance(detailsResponse responses.ManagerDetails, teamsParser helpers.TeamsParser) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	_, err := p.Fprintf(tr, "%s\t%d\n", "Total transfers: ", detailsResponse.LastDeadlineTotalTransfers)

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t£%.1f\n", "Squad value: ", teamsParser.CalculateMonetaryValue(detailsResponse.LastDeadlineValue))

	if err != nil {
		log.Fatal(err)
	}

	_, err = p.Fprintf(tr, "%s\t£%.1f\n", "In the bank: ", teamsParser.CalculateMonetaryValue(detailsResponse.LastDeadlineBank))

	if err != nil {
		log.Fatal(err)
	}

	helpers.AutoFlush(tr)
}

// PrintGameweekFixtures prints the fixtures for the specified gameweek
func (t TerminalRenderer) PrintGameweekFixtures(bootstrap responses.BootstrapData, fixtures responses.GameweekFixtures, teamsParser helpers.TeamsParser, gameweek int) {
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)
	for _, fixture := range fixtures {
		if fixture.Event == gameweek {
			var homeTeam = teamsParser.GetTeam(fixture.HomeTeam, bootstrap)
			var awayTeam = teamsParser.GetTeam(fixture.AwayTeam, bootstrap)

			if fixture.Started {
				var homeScore = fixture.HomeTeamScore
				var awayScore = fixture.AwayTeamScore
				fmt.Fprintf(tr, "%s\t%d-%d\t%s\n", homeTeam, homeScore, awayScore, awayTeam)

			} else {
				fmt.Fprintf(tr, "%s\tvs\t%s\n", homeTeam, awayTeam)
			}
		}
	}

	helpers.AutoFlush(tr)
}

// PrintSeasonDetails prints a teams current season details for each gameweek
func (t TerminalRenderer) PrintSeasonDetails(history responses.ManagerHistory, teamsParser helpers.TeamsParser) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

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

	for _, gameweek := range history.Current {
		_, err := p.Fprintf(tr, "%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%.1f\n",
			gameweek.Event,
			gameweek.Points,
			gameweek.PointsOnBench,
			gameweek.Rank,
			gameweek.EventTransfers,
			gameweek.EventTransfersCost,
			gameweek.TotalPoints,
			gameweek.OverallRank,
			teamsParser.CalculateMonetaryValue(gameweek.Value))

		if err != nil {
			log.Fatal(err)
		}
	}

	helpers.AutoFlush(tr)
}

// PrintChipDetails prints details around what chips have been used
func (t TerminalRenderer) PrintChipDetails(history responses.ManagerHistory) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	fmt.Fprintf(tr, "%s\t%s\t%s\n", "Date", "Name", "Gamweeek used")

	for _, chip := range history.Chips {
		_, err := p.Fprintf(tr, "%s\t%s\t%d\n", chip.Time, chip.Name, chip.Event)

		if err != nil {
			log.Fatal(err)
		}
	}

	helpers.AutoFlush(tr)
}

// PrintPreviousSeasonDetails prints a teams record from past seasons
func (t TerminalRenderer) PrintPreviousSeasonDetails(history responses.ManagerHistory) {
	p := message.NewPrinter(language.English)
	tr := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.FilterHTML)

	fmt.Fprintf(tr, "%s\t%s\t%s\n", "Season", "Points", "Rank")

	for _, season := range history.Past {
		_, err := p.Fprintf(tr, "%s\t%d\t%d\n", season.SeasonName, season.TotalPoints, season.Rank)

		if err != nil {
			log.Fatal(err)
		}
	}

	helpers.AutoFlush(tr)
}

// PrintNoRivalsWarning prints a warning when no rivals have been specified in the config
func (t TerminalRenderer) PrintNoRivalsWarning() {
	fmt.Println("No rivals specified. Update config for this to work.")
}

//go:generate mockgen -source=ui.go -package=ui -destination=mock_renderer.go
