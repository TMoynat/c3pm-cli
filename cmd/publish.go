package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish [project]",
	Short: "Publish your project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage:\n  ctpm publish [string]")
			os.Exit(1)
		}
		fmt.Println("publish called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
