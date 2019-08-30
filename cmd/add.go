package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"ctpm/constants"
)

var addCmd = &cobra.Command{
	Use:   "add [dependency]",
	Short: "Add a dependency to your project",

	RunE: func(cmd *cobra.Command, args []string) error {
		readConfigMandatory()

		if len(args) < 1 {
			return errors.New(constants.MissingCommandArgument)
		}

		newDependency := args[0]
		fmt.Printf("add called with: %s\n", newDependency)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
