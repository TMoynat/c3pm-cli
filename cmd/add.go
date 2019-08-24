package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [dependency]",
	Short: "Add a dependency to your project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
