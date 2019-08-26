package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [script]",
	Short: "Run scripts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
