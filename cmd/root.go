package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"ctpm/constants"
)

var rootCmd = &cobra.Command{
	Use:   "ctpm",
	Short: "ctpm (c3pm) is a package manager for C++",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s.\n", err)
		fmt.Println("Please use the help command to know more.")

		os.Exit(constants.CommandExitStatus)
	}
}

func init() {
	// We silence usage here as the cmdError type
	// invites the user to display the usage when displayed
	rootCmd.SilenceUsage = true

	// We silence the errors as we handle the printing ourselves
	rootCmd.SilenceErrors = true

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// We are looking at "./c3pm.yml"
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("c3pm")

	// Commands are then responsible for reading the configuration file
	// when needed
}
