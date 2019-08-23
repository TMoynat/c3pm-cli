package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unpublishCmd represents the unpublish command
var unpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Unpublish your project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unpublish called")
	},
}

func init() {
	rootCmd.AddCommand(unpublishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unpublishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unpublishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
