package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"ctpm/constants"
)

var addCmd = &cobra.Command{
	Use:   "add [dependency]",
	Short: "Add a dependency to your project",

	RunE: func(cmd *cobra.Command, args []string) error {
		readConfigFatal()

		if len(args) < 1 {
			return errors.New(constants.MissingCommandArgument)
		}

		newDependency := args[0]
		viper.Set("dependencies", newDependency)

		writeConfigFatal()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
