package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [dependency]",
	Short: "Remove a dependency to your project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage:\n  ctpm rm [string]")
			os.Exit(1)
		}
		fmt.Println("rm called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
