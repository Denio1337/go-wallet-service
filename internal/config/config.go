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
	config_flag = "config"

	// Default config file path
	config_default_path = "config.env"

	// Config flag description
	config_flag_description = "path to config file"
)

func init() {
	// Get config path from cmd flag
	configPath := flag.String(config_flag, config_default_path, config_flag_description)
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
