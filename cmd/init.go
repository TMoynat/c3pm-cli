package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		initProject()
	},
}

func initProject() {
	var text string
	fmt.Print("Project name: ")
	fmt.Scanln(&text)
	viper.Set("name", text)
	fmt.Print("Author: ")
	fmt.Scanln(&text)
	viper.Set("author", text)
	fmt.Print("Short description (1 line): ")
	fmt.Scanln(&text)
	viper.Set("description", text)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
