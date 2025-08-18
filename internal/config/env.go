package config

type EnvKey string

// Env params keys
const (
	EnvAppAddress EnvKey = "APP_ADDRESS"
	EnvDBHost     EnvKey = "DB_HOST"
	EnvDBPort     EnvKey = "DB_PORT"
	EnvDBUser     EnvKey = "DB_USER"
	EnvDBPassword EnvKey = "DB_PASSWORD"
	EnvDBName     EnvKey = "DB_NAME"
)

// Check env key validity
func (e EnvKey) isValid() bool {
	switch e {
	case EnvAppAddress, EnvDBHost, EnvDBPort, EnvDBUser,
		EnvDBPassword, EnvDBName:
		return true
	default:
		return false
	}
}
