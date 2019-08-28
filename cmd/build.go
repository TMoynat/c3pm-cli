package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [project]",
	Short: "Build your project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage:\n  ctpm build [string]")
			os.Exit(1)
		}
		fmt.Println("build called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
