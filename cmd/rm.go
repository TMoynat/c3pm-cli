package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [dependency]",
	Short: "Remove a dependency to your project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
