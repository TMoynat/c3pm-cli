package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"ctpm/constants"
)

func readConfigFatal() {
	// Read configuration, success or die
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf(
			"Could not read %s. Have you created a c3pm project?\n",
			constants.ConfigurationFileName)
		os.Exit(constants.ConfigurationReadExitStatus)
	}
}

func writeConfigFatal() {
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("Error: %s.\n", err)
		fmt.Println("Could not write configuration.")

		os.Exit(constants.ConfigurationWriteExitStatus)
	}
}
