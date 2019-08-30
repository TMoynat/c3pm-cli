package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [dependency]",
	Short: "Remove a dependency to your project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Usage:\n  ctpm remove [string]")
		}
		fmt.Println("rm called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
