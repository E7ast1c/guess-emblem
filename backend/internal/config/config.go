package config

type appConfig struct {}

type AppConfig interface {
	Value(key string) interface{}
	All() FullConfig
}

type FullConfig struct {
	DBConfig  DBConfig
	APIServer APIServer
}

type DBConfig struct {
	URI string `env:"DB_URI"`
}

type APIServer struct {
	Port       string `env:"PORT"`
	SignSecret string `env:"SIGN_SECRET"`
}

func New() AppConfig {
	return &appConfig{}
}

func (c *appConfig) All() FullConfig {
	return FullConfig{
		DBConfig:  DBConfig{URI: MustEnvString("DB_URI")},
		APIServer: APIServer{
			Port:       MustEnvString("PORT"),
			SignSecret: MustEnvString("SIGN_SECRET"),
		},
	}
}

func (c *appConfig) Value(key string) interface{} {
	return MustEnvString(key)
}
