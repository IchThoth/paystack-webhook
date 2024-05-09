package config

import (
	"os"
	"strconv"
)

func(a *AppConfig) env() string {
	const env = "production"

	if newEnv, exists := os.LookupEnv("ENV"); exists {
		return newEnv
	}

	return env
}

func (a *AppConfig) port() int {
	const port = 8045

	if newPort, exists := os.LookupEnv("PORT"); exists {
		intPort, err := strconv.Atoi(newPort)
		if err == nil {
			return intPort
		}
	}

	return port
}


func (a *AppConfig) dbDsn() string {
	const dbDsn =  ""

	if dsn , exists := os.LookupEnv("DB_DSN"); exists {
		return dsn
	}

	return dbDsn
}

func (a *AppConfig) dbMaxOpenConn() int   {
	const maxconns = 25

	if maxOpen , exists := os.LookupEnv("DB_MAX_OPEN");  exists {
		nxMaxOpen, err := strconv.Atoi(maxOpen)
		if err == nil {
			return nxMaxOpen
		}
	}

	return maxconns
}

func (a *AppConfig) dbMaxIdle() string {
	const maxidle = "15"

	if maxIdle, exists :=os.LookupEnv("DB_MAX_IDLE"); exists {
		return maxIdle
	}
	return maxidle
}

func (a *AppConfig) mailHost() string {
	if host, exists := os.LookupEnv("MAIL_HOST"); exists {
		return host
	}
	return ""
}

func (a *AppConfig) mailPort() int {
	const port = 455

	if mport,exists:= os.LookupEnv("MAIL_PORT"); exists {
		intPort, err := strconv.Atoi(mport)
		if err == nil {
			return intPort
		}
	}

	return port
}

func (a *AppConfig) mailUser() string  {
	if user, exists := os.LookupEnv("MAIL_USER"); exists {
		return user
	}

	return ""
}

