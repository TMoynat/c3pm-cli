package cmd

import (
	"fmt"
	"errors"
	"os"

	"github.com/spf13/viper"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"ctpm/constants"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new c3pm project",

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("User wants to initialize a project. This part will be interactive.")

		viper.Set("name", "this_is_a_test")
		_ = viper.WriteConfigAs(constants.ConfigurationFileName)

		initProject()
		return nil
	},
}

func initProject() {
	validateN := func(input string) error {
		if len(input) < 1 {
			return errors.New("Invalid name")
		}
		return nil
	}

	validateP := func(input string) error {
		err := os.Mkdir(input, 0755)
		if err != nil {
			return errors.New("Invalid path")
		}
		return nil
	}

	name := promptui.Prompt{
		Label:    "Project name ",
		Validate: validateN,
	}

	pName, err := name.Run()
	checkError(err)

	author := promptui.Prompt{
		Label:    "Author ",
	}

	pAuth, err := author.Run()
	checkError(err)

	path := promptui.Prompt{
		Label:    "Project location ",
		Validate: validateP,
	}

	pPath, err := path.Run()
	checkError(err)

	fmt.Println("You choose : ", pName, pAuth, pPath)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(constants.CommandExitStatus)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
