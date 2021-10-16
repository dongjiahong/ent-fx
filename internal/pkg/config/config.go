package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config interface {
	GetDBConfig() *DBConfig
	GetWebConfig() *WebConfig
}

type DBConfig struct {
	User        string `toml:"User"`
	Password    string `toml:"Password"`
	Host        string `toml:"Host"`
	DBName      string `toml:"DBName"`
	TablePrefix string `toml:"TablePrefix"`
}

type WebConfig struct {
	EndPoint int `toml:"EndPoint"`
}

type config struct {
	DBConfig  *DBConfig  `toml:"DBConfig"`
	WebConfig *WebConfig `toml:"WebConfig"`
}

type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (f optionFunc) apply(c *config) { f(c) }

func NewConfig(opts ...Option) Config {
	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) == 0 {
		panic("can't find config path")
	}

	var instance config
	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigFile(configPath)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&instance); err != nil {
		panic(err)
	}

	for _, opt := range opts {
		opt.apply(&instance)
	}
	return &instance
}

func (c *config) GetDBConfig() *DBConfig {
	return c.DBConfig
}

func (c *config) GetWebConfig() *WebConfig {
	return c.WebConfig
}
