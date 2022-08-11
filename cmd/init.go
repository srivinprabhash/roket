/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Connection details struct
type Connection struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// roket.yaml struct
type RoketYaml struct {
	Connection Connection `yaml:"connection"`
	Migrations string     `yaml:"migrations"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes roket.",
	Long:  `Initializes roket in the current directory. Creates a roket.yaml file.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			fmt.Println(`
Migrations path is not specified. Default
path will be used. (database/migrations) 
 `)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		isForced, _ := cmd.Flags().GetBool("force")
		var path string

		if len(args) > 0 {
			path = args[0]
		} else {
			path = "database/migrations"
		}

		// Check if given migrations path exist
		_, err := os.Stat(path)
		if !os.IsExist(err) {

			// Check if forced option is provided
			if !isForced {

				prompt := promptui.Select{
					Label: "Migrations path already exists. Do you want to overwrite it's content ?",
					Items: []string{
						"Yes",
						"No",
					},
					HideHelp: true,
				}

				_, r, err := prompt.Run()
				if err != nil {
					log.Fatalln(err)
				}

				if r != "Yes" {
					return
				}

			}

		}

		// Create migrations path
		err2 := os.MkdirAll(path, 0770)
		if err2 != nil {
			log.Fatalln(err2)
		}

		// Init roket.yaml
		// Check if roket.yaml already exist
		if !isForced {

			_, err := os.Stat("roket.yaml")
			if !os.IsExist(err) {

				prompt := promptui.Select{
					Label: "roket.yaml file already exists. Do you want to overwrite it ?",
					Items: []string{
						"Yes",
						"No",
					},
					HideHelp: true,
				}

				_, r, err := prompt.Run()
				if err != nil {
					log.Fatalln(err)
				}

				if r != "Yes" {
					fmt.Println("Aborting roket init")
					return
				}

			}

		}

	},
}

func initRoket(path string, args []string) (*RoketYaml, error) {

	// Create roket.yaml
	config := RoketYaml{
		Migrations: "database/migrations",
		Connection: Connection{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     3306,
			Database: "database",
			User:     "user",
			Password: "password",
		},
	}

	config.Migrations = path

	yamlFile, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

	// Writing roket.yaml
	err2 := os.WriteFile("roket.yaml", yamlFile, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
	return &config, nil

}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("force", "f", false, "Force the init command. Will replace a roket.yaml file if alraedy exist.")

}
