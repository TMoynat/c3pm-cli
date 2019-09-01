package cmd

import (
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

type Config struct {
	name, author, description, version string
}

var providedFolder string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new c3pm project",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			err := initProject()
			if err != nil {
				return constants.ConfigurationWriteExitStatus
			}
		} else {
			if _, err := os.Stat(args[0]); !os.IsNotExist(err) {
				return errors.New("Invalid path")
			}
			os.MkdirAll(args[0], 0755)
			providedFolder = args[0]
			err := initProject()
			if err != nil {
				return constants.ConfigurationWriteExitStatus
			}
		}
		return nil
	},
}

func initProject() error {
	validateVersion := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid version")
		}
		return nil
	}

	var project Config
	name := promptui.Prompt{
		Label:    "Project name ",
	}
	pName, err := name.Run()
	if err != nil {
		return err
	} else {
		project.name = pName
	}

	author := promptui.Prompt{
		Label:    "Author ",
	}
	pAuth, err := author.Run()
	if err != nil {
		return err
	} else {
		project.author = pAuth
	}

	version := promptui.Prompt{
		Label:    "Version ",
		Validate: validateVersion,
	}
	pVers, err := version.Run()
	if err != nil {
		return err
	} else {
		project.version = pVers
	}

	description := promptui.Prompt{
		Label:    "Description ",
	}
	pDesc, err := description.Run()
	if err != nil {
		return err
	} else {
		project.description = pDesc
	}
	return configProject(project)
}

func configProject(project Config) error {
	if len(project.name) < 1 {
		dir, err := os.Getwd()
		viper.Set("name", filepath.Base(dir))
		if err != nil {
			return err
		}
	} else {
		viper.Set("name", project.name)
	}
	viper.Set("author", project.author)
	if len(project.version) < 3 {
		viper.Set("version", "1.0.0")
	} else {
		s := []string{project.version, ".0"}
		viper.Set("version", strings.Join(s, ""))
	}
	viper.Set("description", project.description)
	err := viper.WriteConfigAs(filepath.Join(providedFolder, constants.ConfigurationFileName))
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
