package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiate a new c3pm project",

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("User wants to initiate a project. This part will be interactive.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
