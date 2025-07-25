package config

import (
	"log"

	"github.com/shanth1/gotools/conf"
)

func MustGetConfig() *Config {
	cfg := &Config{}
	if err := conf.Load(conf.GetConfigPath(), cfg); err != nil {
		log.Fatalf("load config: %v", err)
	}

	return cfg
}

type Config struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	UseMocks      bool   `mapstructure:"USE_MOCKS"`
	Neo4jURI      string `mapstructure:"NEO4J_URI"`
	Neo4jUser     string `mapstructure:"NEO4J_USER"`
	Neo4jPassword string `mapstructure:"NEO4J_PASSWORD"`
	RedisAddr     string `mapstructure:"REDIS_ADDR"`
}
