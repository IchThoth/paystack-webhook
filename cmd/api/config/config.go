package config

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

func (a *AppConfig) Parse() {

}
