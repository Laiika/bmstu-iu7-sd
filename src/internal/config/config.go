package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sd/pkg/logger"
	"sync"
)

type Config struct {
	Postgres      PostgresConfig `yaml:"postgres"`
	Mongo         MongoConfig    `yaml:"mongodb"`
	TelegramToken string         `yaml:"telegram_token"`
}

type PostgresConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port" default:"5432"`
	Database string `yaml:"dbname"`
}

type MongoConfig struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port" default:"27017"`
	Database   string `yaml:"dbname"`
	AuthSource string `yaml:"auth_source"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logger.Logger) *Config {
	once.Do(func() {
		logger.Info("чтение конфигурационного файла")
		instance = &Config{}
		if err := cleanenv.ReadConfig("./config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
