package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	PORT int `yaml:"port" env-required:"true"`
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server" env-required:"true"`
}

func Load() *Config {
	var configPath string
    configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
        configPath = "config/config.yaml"
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath,&cfg)
	if err != nil {
       log.Fatal("config file is not valid")
	}

	return &cfg
}


