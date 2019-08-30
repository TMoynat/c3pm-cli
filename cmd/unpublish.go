package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// unpublishCmd represents the unpublish command
var unpublishCmd = &cobra.Command{
	Use:   "unpublish [project]",
	Short: "Unpublish your project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Usage:\n  ctpm unpublish [string]")
		}
		fmt.Println("unpublish called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(unpublishCmd)
}
