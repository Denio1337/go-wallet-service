package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	// Config flag name
	configFlag = "config"

	// Default config file path
	configPath = "config.env"

	// Config flag description
	configFlagDescr = "path to config file"
)

type Config struct {
	AppConfig
	StorageConfig
}

type AppConfig struct {
	Address string `env:"APP_ADDRESS"`
}

type StorageConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

func MustLoad() *Config {
	// Get config path from cmd flag
	configPath := flag.String(configFlag, configPath, configFlagDescr)
	flag.Parse()

	// Check if file exists
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", *configPath)
	}

	var cfg Config

	// Read config from file
	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	return &cfg
}
