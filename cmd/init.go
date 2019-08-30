package cmd

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"ctpm/constants"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new c3pm project",

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("User wants to initialize a project. This part will be interactive.")

		viper.Set("name", "this_is_a_test")
		_ = viper.WriteConfigAs(constants.ConfigurationFileName)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
