package datastore

import (
	config "awawe/configuration"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", config.GetMySQLConfig().DSN)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(config.GetMySQLConfig().MaxOpen)
	db.SetMaxIdleConns(config.GetMySQLConfig().MaxIdle)

	return db
}
