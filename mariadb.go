package mariadb

import (
	"database/sql"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
)

type MariaDB struct {
	*sql.DB
}

var _ database.SQLExecutor = MariaDB{}
var _ database.SQLQuerier = MariaDB{}
var _ database.SQLPreparation = MariaDB{}

func New(conf config.Database) (*MariaDB, error) {
	db, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		return nil, err
	}

	return &MariaDB{db}, nil
}
