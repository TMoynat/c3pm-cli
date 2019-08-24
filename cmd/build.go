package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [project]",
	Short: "Build your project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
