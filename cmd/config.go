package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"ctpm/constants"
)

func readConfigMandatory() {
	// Read configuration, success or die
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf(
			"Could not read %s. Have you created a c3pm project?\n",
			constants.ConfigurationFileName)
		os.Exit(constants.ConfigurationError)
	}
}
