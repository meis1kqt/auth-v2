package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env"`
	HTTPServer `yaml:"http-server"`
	DataBase   `yaml:"database"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_SERVER_ADDRESS" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_SERVER_TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"10s"`
}

type DataBase struct {
	Host     string `yaml:"host" env:"DATABASE_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DATABASE_PORT" env-default:"5432"`
	User     string `yaml:"user" env:"DATABASE_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD" env-default:"postgres"`
	DBName   string `yaml:"dbname" env:"DATABASE_DBNAME" env-default:"auth_db"`
}


func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "./config/local.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("env invalid")
	}

	return &cfg


}