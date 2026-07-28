package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host            string
	Port            string
	Name            string
	User            string
	Password        string
	Enable_ssl_mode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

var configurations *Config

func loadConfig() {
	// Load .env file
	// err := godotenv.Load()
	err := godotenv.Overload();
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// ========================
	// Application Configuration
	// ========================

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME is required")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT_SECRET_KEY is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP_PORT is required")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("HTTP_PORT must be a valid number")
		os.Exit(1)
	}

	// ========================
	// Database Configuration
	// ========================

	dbHost := os.Getenv("HOST")
	if dbHost == "" {
		fmt.Println("HOST is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("PORT")
	if dbPort == "" {
		fmt.Println("PORT is required")
		os.Exit(1)
	}

	_, err = strconv.Atoi(dbPort)
	if err != nil {
		fmt.Println("PORT must be a valid number")
		os.Exit(1)
	}

	dbName := os.Getenv("NAME")
	if dbName == "" {
		fmt.Println("NAME is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("USER")
	if dbUser == "" {
		fmt.Println("USER is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("PASSWORD")
	if dbPassword == "" {
		fmt.Println("PASSWORD is required")
		os.Exit(1)
	}

	dbEnableSSLMode := os.Getenv("ENABLE_SSL_MODE")
	if dbEnableSSLMode == "" {
		dbEnableSSLMode = "false"
	}

	enableSSLMode, err := strconv.ParseBool(dbEnableSSLMode)
	if err != nil {
		fmt.Println("ENABLE_SSL_MODE must be true or false")
		os.Exit(1)
	}

	// Create DB Config
	dbConfiguration := &DBConfig{
		Host:            dbHost,
		Port:            dbPort,
		Name:            dbName,
		User:            dbUser,
		Password:        dbPassword,
		Enable_ssl_mode: enableSSLMode,
	}

	// Create Main Config
	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfiguration,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}
	return configurations
}
