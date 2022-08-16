/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new migration.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Migration name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		// Concatname migration dir name
		path := viper.GetViper().GetString("migrations")
		unixtime := strconv.FormatInt(time.Now().Unix(), 10)
		migrationName := unixtime + "_" + args[0]
		path = path + "/" + migrationName

		// Create migration dir
		if err := os.Mkdir(path, 0744); err != nil {
			fmt.Println(err)
			return
		}

		// Create up/down migration files
		if _, err := os.Create(path + "/up.sql"); err != nil {
			fmt.Println(colorRed, "Error :: ", colorReset, "Could not create up migration file.")
			return
		}

		if _, err := os.Create(path + "/down.sql"); err != nil {
			fmt.Println(colorRed, "Error :: ", colorReset, "Could not create down migration file.")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
