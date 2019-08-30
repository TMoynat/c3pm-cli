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

		initProject()
		return nil
	},
}

func initProject() {
	validateName := func(input string) error {
		if len(input) < 1 {
			return errors.New("Invalid name")
		}
		return nil
	}

	validatePath := func(input string) error {
		if _, err := os.Stat(input); !os.IsNotExist(err) {
			return errors.New("Invalid path")
		}
		return nil
	}

	name := promptui.Prompt{
		Label:    "Project name ",
		Validate: validateName,
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
		Validate: validatePath,
	}
	pPath, err := path.Run()
	checkError(err)
	os.Mkdir(pPath, 0755)

	description := promptui.Prompt{
		Label:    "Description ",
	}
	pDesc, err := description.Run()
	checkError(err)

	configProject(pName, pAuth, pDesc)
}

func configProject(pName, pAuth, pDesc string) {
	viper.Set("name", pName)
	viper.Set("author", pAuth)
	viper.Set("description", pDesc)
	err := viper.WriteConfig()
	checkError(err)
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
