package connection

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		DBUser:     os.Getenv("DBUSER"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBHost:     os.Getenv("DBHOST"),
		DBPort:     os.Getenv("DBPORT"),
		DBName:     os.Getenv("DBNAME"),
	}
	fields := map[string]string{
		"DBUser":     cfg.DBUser,
		"DBPassword": cfg.DBPassword,
		"DBHost":     cfg.DBHost,
		"DBPort":     cfg.DBPort,
		"DBName":     cfg.DBName,
	}

	for name, val := range fields {
		if val == "" {
			return nil, fmt.Errorf("missing required database configuration: %s", name)
		}
	}

	return cfg, nil
}
