package config

import (
	"fmt"
	"os"
	"strconv"
	"transaction/pkg/log"
)

var logger = log.NewLogger(os.Stdout)

func getStrEnv(key string, fallback string, required bool) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	if required {
		logger.Fatal("Missing environment variable ", key)
	}
	return fallback
}

func getIntEnv(key string, fallback int, required bool) int {

	if valueStr, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			logger.Fatalf("Invalid variable ", key, "with value", valueStr)
		}
		return value
	}
	if required {
		logger.Fatal("Variable ", key, "is required")
		return 0
	} else {
		return fallback
	}
}

func getBoolEnv(key string, fallback bool, required bool) bool {

	if valueStr, ok := os.LookupEnv(key); ok {
		if valueStr == "1" || valueStr == "true" {
			return true
		} else if valueStr == "0" || valueStr == "false" {
			return false
		} else {
			logger.Fatalf("Invalid variable ", key)
			return false
		}
	}
	if required {
		logger.Fatal("Variable ", key, "is required")
		return false
	} else {
		return fallback
	}
}

type DBConfig struct {
	Name       string
	DriverName string
	Host       string
	Port       int
	User       string
	Password   string
}

type Config struct {
	DB    *DBConfig
	Port  int
	Debug bool
}

var Default = Config{}

func (c *DBConfig) DSN() string {
	fmt.Printf("%s://%s:%d/%s?sslmode=disable&user=%s&password=%s",
		c.DriverName,
		c.Host,
		c.Port,
		c.Name,
		c.User,
		c.Password)
	return fmt.Sprintf("%s://%s:%d/%s?sslmode=disable&user=%s&password=%s",
		c.DriverName,
		c.Host,
		c.Port,
		c.Name,
		c.User,
		c.Password,
	)
}
func (c *DBConfig) ConnString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func GetConfig() *Config {

	dbConfig := &DBConfig{
		Name:       getStrEnv("DB_NAME", "", true),
		DriverName: getStrEnv("DRIVER_NAME", "pq", false),
		Host:       getStrEnv("DB_HOST", "localhost", false),
		Port:       getIntEnv("DB_PORT", 5432, false),
		User:       getStrEnv("DB_USER", "", true),
		Password:   getStrEnv("DB_PASSWORD", "", true),
	}

	return &Config{
		DB:    dbConfig,
		Port:  getIntEnv("PORT", 8080, false),
		Debug: getBoolEnv("DEBUG", false, false),
	}
}

func LoadFromDotEnv(path string) *Config {
	return nil
}
