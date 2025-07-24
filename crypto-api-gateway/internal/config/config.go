package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPC *GRPCConfig
	HTTP *HttpConfig
}

type HttpConfig struct {
	Host string
	Port string
}

type GRPCConfig struct {
	OrderHost string
	OrderPort string
	AssetHost string
	AssetPort string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		HTTP: loadHttpConfig(),
		GRPC: loadGRPCConfig(),
	}, nil
}

func loadGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		OrderHost: getEnv("GRPC_ORDER_HOST"),
		OrderPort: getEnv("GRPC_ORDER_PORT"),
		AssetHost: getEnv("GRPC_ASSET_HOST"),
		AssetPort: getEnv("GRPC_ASSET_PORT"),
	}
}

func loadHttpConfig() *HttpConfig {
	return &HttpConfig{
		Host: getEnv("HTTP_HOST"),
		Port: getEnv("HTTP_PORT"),
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
