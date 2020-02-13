package configuration

const (
	ENV = "AWAWE_ENV"
)

type (
	app struct {
		Command        string `json:"command"`
		AutoMigrate    bool   `json:"auto_migrate"`
		MigrationSteps int    `json:"migration_step"`
		MigrationPath  string `json:"migration_path"`
	}
	server struct {
		Host         string `json:"host"`
		Port         string `json:"port"`
		ReadTimeout  int    `json:"read_timeout"`
		WriteTimeout int    `json:"write_timeout"`
	}
	mySQLConfig struct {
		DSN     string `json:"dsn"`
		MaxIdle int    `json:"max_idle"`
		MaxOpen int    `json:"max_open"`
	}
	logger struct {
		Level    string `json:"level"`
		Output   string `json:"output"`
		Filename string `json:"filename"`
	}
	appRedis struct {
		Address      string `json:"add"`
		Password     string `json:"password"`
		DB           int    `json:"db"`
		ReadTimeout  int    `json:"read_timeout"`
		WriteTimeout int    `json:"write_timeout"`
		PoolSize     int    `json:"pool_size"`
	}
)

var config = &struct {
	App         *app         `json:"app"`
	Server      *server      `json:"server"`
	MySQLConfig *mySQLConfig `json:"mysql_config"`
	Logger      *logger      `json:"logger"`
	AppRedis    *appRedis    `json:"app_redis"`
}{}

func GetAppConfig() *app {
	return config.App
}

func GetServerConfig() *server {
	return config.Server
}

func GetMySQLConfig() *mySQLConfig {
	return config.MySQLConfig
}

func GetLoggerConfig() *logger {
	return config.Logger
}

func GetRedisConfig() *appRedis {
	return config.AppRedis
}

func InitializeConfig() {
	loadJSONEnvPathOrPanic(ENV, config)
}
