package cmd

import (
	"fmt"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	validatePath := func(input string) error {
		if _, err := os.Stat(input); !os.IsNotExist(err) {
			return errors.New("Invalid path")
		}
		return nil
	}

	validateVersion := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid version")
		}
		return nil
	}

	name := promptui.Prompt{
		Label:    "Project name ",
	}
	pName, err := name.Run()
	checkError(err)

	author := promptui.Prompt{
		Label:    "Author ",
	}
	pAuth, err := author.Run()
	checkError(err)

	version := promptui.Prompt{
		Label:    "Version ",
		Validate: validateVersion,
	}
	pVers, err := version.Run()
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

	configProject(pName, pAuth, pDesc, pVers)
}

func configProject(pName, pAuth, pDesc, pVers string) {
	if len(pName) < 1 {
		dir, err := os.Getwd()
		viper.Set("name", filepath.Base(dir))
		checkError(err)
	} else {
		viper.Set("name", pName)
	}
	viper.Set("author", pAuth)
	if len(pVers) < 3 {
		viper.Set("version", "1.0.0")
	} else {
		s := []string{pVers, ".0"}
		viper.Set("version", strings.Join(s, ""))
	}
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
