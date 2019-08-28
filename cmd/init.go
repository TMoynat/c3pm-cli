package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1{
			fmt.Println("Usage:\n  ctpm init [string]")
			os.Exit(1)
		}
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
