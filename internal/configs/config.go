package configs

import "os"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() error {
	cfg = &config{
		API: APIConfig{
			Port: getEnv("API_PORT", "8080"),
		},
		DB: DBConfig{
			Host:     getEnv("DATABASE_HOST", "localhost"),
			Port:     getEnv("DATABASE_PORT", "5432"),
			User:     getEnv("DATABASE_USER", ""),
			Password: getEnv("DATABASE_PASSWORD", ""),
			Database: getEnv("DATABASE_NAME", ""),
		},
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetDb() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}