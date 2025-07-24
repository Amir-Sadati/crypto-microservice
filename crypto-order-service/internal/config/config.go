package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Postgres *PostgresConfig
	GRPC     *GRPCConfig
}

type PostgresConfig struct {
	User     string
	Password string
	DBName   string
	Port     int
	Url      string
}

type GRPCConfig struct {
	Host string
	Port string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Postgres: loadPostgresConfig(),
		GRPC:     loadGRPCConfig(),
	}, nil
}

func loadPostgresConfig() *PostgresConfig {

	return &PostgresConfig{
		User:     getEnv("POSTGRES_USER"),
		Password: getEnv("POSTGRES_PASSWORD"),
		DBName:   getEnv("POSTGRES_DB"),
		Port:     getEnvAsInt("POSTGRES_PORT"),
		Url:      getEnv("POSTGRES_URL"),
	}
}

func loadGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		Host: getEnv("GRPC_HOST"),
		Port: getEnv("GRPC_PORT"),
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing env var: %s", key)
	}
	return val
}

func getEnvAsInt(key string) int {
	val := getEnv(key)
	n, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Invalid int for %s: %v", key, err)
	}
	return n
}
