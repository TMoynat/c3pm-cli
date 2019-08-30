package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [script]",
	Short: "Run scripts",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Usage:\n  ctpm run [string]")
		}
		fmt.Println("run called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
