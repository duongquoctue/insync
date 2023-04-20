package mysql

import (
	"database/sql"
	sqldriver "github.com/go-sql-driver/mysql"
)

const (
	maxOpenConnection int = 100
	maxIdleConnection int = 10
)

func NewConnection(cfg *sqldriver.Config) (err error) {
	if cfg == nil {
		cfg, err = getLocalConfig()
	}

	if err != nil {
		return
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return
	}

	db.SetMaxOpenConns(maxOpenConnection)
	db.SetMaxIdleConns(maxIdleConnection)

	err = db.Ping()
	return
}
