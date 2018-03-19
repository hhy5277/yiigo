package yiigo

// X is a convenient alias for a map[string]interface{} map
type X map[string]interface{}

// Bootstrap init and start components
func Bootstrap(mysql bool, mongo bool, redis bool) error {
	// load config
	loadEnv("env.toml")
	// init logger
	initLogger()

	if mysql {
		// init mysql
		if err := initMySQL(); err != nil {
			return err
		}
	}

	if mongo {
		// init mongodb
		if err := initMongo(); err != nil {
			return err
		}
	}

	if redis {
		// init redis
		if err := initRedis(); err != nil {
			return err
		}
	}

	return nil
}
