package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"path"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
}

func (c *Config) isValid() bool {
	return c.DB.Host != "" &&
		c.DB.Port != "" &&
		c.DB.User != "" &&
		c.DB.Password != "" &&
		c.DB.DBName != "" &&
		c.Server.Address != ""
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Address string
}

func Load() (*Config, error) {

	pathEnv := path.Join("D:/ApplicationBackend/.env")
	err := godotenv.Load(pathEnv)

	if err != nil {
		return nil, fmt.Errorf("No .env file found, using system environment variables")
	}

	config := &Config{
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
		Server: ServerConfig{
			Address: os.Getenv("SERVER_PORT"),
		},
	}

	if config.isValid() {
		return config, nil

	} else {
		return nil, fmt.Errorf("empty config")
	}
}
