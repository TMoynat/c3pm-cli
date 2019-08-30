package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [dependency]",
	Short: "Add a dependency to your project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Usage: ctpm add [string]")
		}
		viper.Set("Dependencies", args[0])
		fmt.Println("add called with arg :", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
