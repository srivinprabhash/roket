package connections

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func Connect() (*sqlx.DB, error) {

	viper := viper.GetViper()

	if viper.GetString("connection.driver") == "mysql" {

		connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s", viper.GetString("connection.user"), viper.GetString("connection.password"), viper.GetString("connection.host"), viper.GetString("connection.port"), viper.GetString("connection.database"))
		db, err := sqlx.Connect("mysql", connStr)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		return db, nil
	}

	return nil, errors.New("connection driver not defined")

}
