package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [dependency]",
	Short: "Remove a dependency to your project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Usage: ctpm remove [string]")
		}
		fmt.Println("remove called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
