package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Flags
const (
	// Config flag name
	configFlag = "config"

	// Default config file path
	configPath = "config.env"

	// Config flag description
	configFlagDescr = "path to config file"
)

func init() {
	// Get config path from cmd flag
	configPath := flag.String(configFlag, configPath, configFlagDescr)
	flag.Parse()

	// Load .env file
	err := godotenv.Load(*configPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// Get .env value by key
func Get(key EnvKey) string {
	// Return default value "" if env key is invalid
	if !key.isValid() {
		return ""
	}

	return os.Getenv(string(key))
}

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
	DBName   string `env:"DB_NAME"`
}
