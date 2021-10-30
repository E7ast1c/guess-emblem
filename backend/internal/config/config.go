package config

import "guess-emblem/pkg/env"

type currentConfig struct{}

type ConfigMethods interface {
	GetValues() AppConfig
}

type AppConfig struct {
	DBConfig  DBConfig
	APIServer APIServer
}

type DBConfig struct {
	URI string `env:"DB_URI"`
}

type APIServer struct {
	Port       string `env:"PORT"`
	Address    string `env:"Address"`
	SignSecret string `env:"SIGN_SECRET"`
}

func NewAppConfig() ConfigMethods {
	return &currentConfig{}
}

func (c *currentConfig) GetValues() AppConfig {
	return AppConfig{
		DBConfig: DBConfig{URI: env.MustEnvString("DB_URI")},
		APIServer: APIServer{
			Port:       env.MustEnvString("PORT"),
			SignSecret: env.MustEnvString("SIGN_SECRET"),
		},
	}
}
