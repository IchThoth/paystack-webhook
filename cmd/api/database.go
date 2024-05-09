package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ichthoth/clean-paystack-webhook/cmd/api/config"
)

func openDB(cfg config.AppConfig) (*sql.DB, error) {
	dbconfig := cfg.Db

	db, err := sql.Open("mysql", dbconfig.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(dbconfig.MaxIdleConns)
	db.SetMaxOpenConns(dbconfig.MaxOpenConns)
	
	idletime, err := time.ParseDuration(dbconfig.MaxIdleTimes)

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
