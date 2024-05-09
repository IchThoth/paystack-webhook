package config

import (
	"flag"
)

type AppConfig struct {
	Env     string
	Version string
	Port    int

	Db struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTimes int
	}

	Smtp struct {
		Host     string
		Port     string
		User     string
		Password string
		Sender   string
	}

	paystack struct {
		//access and secret keys will both be test keys in development environment
		access string
		secret string
	}
	//webhook config should be somewhere here
	// will do adeuate research on it

}

var cfg AppConfig

func (a *AppConfig) Parse() {
	flag.StringVar(&cfg.Env, "application environment", a.env(), "")

	//db
	flag.StringVar(&cfg.Db.Dsn, "PostgreSQL Database",a.dbDsn(),"")
	flag.IntVar(&cfg.Db.MaxOpenConns,"Maximum amount of connnections in pool",a.dbMaxOpenConn(),"")
	flag.IntVar(&cfg.Db.MaxIdleConns,"Maximum amount of Idle Connections",a.dbMaxIdle(),"")
	flag.
	//mailer service

	//paystack webhook service.. this will probably exposes payment inforation of users that triggered the
	//webhook
	
	
}

