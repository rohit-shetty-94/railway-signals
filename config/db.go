package config

import (
	"github.com/go-pg/pg/v10"
)

func ConnectDB() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "Rohit@29",
		Database: "railway_signals",
	})
}
