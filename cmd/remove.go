package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [dependency]",
	Short: "Remove a dependency to your project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
