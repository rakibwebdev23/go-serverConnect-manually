package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	ServiceName string
	HttpPort int
}

var configurations Config

func loadConfig() {
	err := godotenv.Load()
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("HTTP port must be a valid number")
		os.Exit(1)
	}

	configurations = Config {
		Version: version,
		ServiceName: serviceName,
		HttpPort: int(port), // type casting
	}

}

func GetConfig () Config {
	loadConfig()
	return configurations
}