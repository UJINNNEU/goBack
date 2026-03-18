package config

import (
    _"fmt"
    _"os"
    _"strconv"
   _ "github.com/joho/godotenv"
)
type Config struct {
    config
}

type config struct {
    DB struct {
        Host     string
        Port     int
        User     string
        Password string
        DBName   string
        SSLMode  string
    }
    Server struct {
        Port string
    }
}

func Load() (*Config, error) {


    
    cfg := &Config{}
    
    cfg.DB.Host = "localhost"
    cfg.DB.Port = 5432
    cfg.DB.User = "postgres"
    cfg.DB.Password = "password"
    cfg.DB.DBName = "your_db_name"
    cfg.DB.SSLMode = "disable"
    
    cfg.Server.Port = ":8080"
    
    return cfg, nil
}