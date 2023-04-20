package mysql

import (
	"time"

	sqldriver "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const (
	dialTimeout = time.Second * 20
)

// TODO: move to consul later
func getLocalConfig() (*sqldriver.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &sqldriver.Config{
		User:    viper.GetString("user"),
		Passwd:  viper.GetString("password"),
		Net:     viper.GetString("network_type"),
		Addr:    viper.GetString("address"),
		DBName:  viper.GetString("database"),
		Timeout: dialTimeout,
	}, nil
}
