/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/srivinprabhash/roket/connections"
)

// Command implementation
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Run: func(cmd *cobra.Command, args []string) {

		// List migrations
		path := viper.GetViper().GetString("migrations")
		migrationDirs, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
		}

		// Preparing migrations
		var m = make(map[int64]string)
		var keys []int64

		for _, each := range migrationDirs {
			k, err := strconv.ParseInt(strings.Split(each.Name(), "_")[0], 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			v := each.Name()
			m[k] = v
			keys = append(keys, k)
		}

		// Sort migrations in ascending order
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})

		// Get connection
		db, err := connections.Connect()
		if err != nil {
			fmt.Println(err)
		}

		for _, v := range keys {

			// Run each migration
			fmt.Println("Running migration ::", m[v])
			migration_path := path + "/" + m[v] + "/up.sql"

			body, err := os.ReadFile(migration_path)
			if err != nil {
				fmt.Println(err)
			}

			_, err = db.Exec(string(body))
			if err != nil {

				var mErr *mysql.MySQLError
				if errors.As(err, &mErr) {
					fmt.Println(mErr.Message)
					return
				} else {
					fmt.Println("Oops ! Something went wrong.")
				}
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
