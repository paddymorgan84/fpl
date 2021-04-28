package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/paddymorgan84/fpl/api"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fpl",
	Short: "A CLI tool for retrieving FPL data",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fpl.yaml)")
	rootCmd.PersistentFlags().StringP("gameweek", "g", "", "The gameweek you wish to see details for")
	err := viper.BindPFlag("gameweek", rootCmd.PersistentFlags().Lookup("gameweek"))
	var fplClient api.FplAPI = api.New()

	if err != nil {
		log.Fatal(err)
	}

	// Add all the commands we want
	rootCmd.AddCommand(BuildDetailsCommand(&fplClient))
	rootCmd.AddCommand(BuildFixturesCommand(&fplClient))
	rootCmd.AddCommand(BuildHistoryCommand(&fplClient))
	rootCmd.AddCommand(BuildPointsCommand(&fplClient))
	rootCmd.AddCommand(BuildRivalsCommand(&fplClient))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".fpl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".fpl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
