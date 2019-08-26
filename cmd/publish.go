package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish [project]",
	Short: "Publish your project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publish called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
