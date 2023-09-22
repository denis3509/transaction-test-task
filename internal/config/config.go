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
	strValue := getStrEnv(key, strconv.Itoa(fallback), required)
	value, err := strconv.Atoi(strValue)
	if err != nil {
		logger.Fatalf("Invalid variable ", key)
	}
	return value
}

func getBoolEnv(key string, fallback bool, required bool) bool {
	var fallbackStr string
	if fallback == false {
		fallbackStr = "0"
	} else {
		fallbackStr = "1"
	}
	strValue := getStrEnv(key, fallbackStr, required)

	if strValue == "1" {
		return true
	} else if strValue == "0" {
		return false
	} else {
		logger.Fatalf("Invalid variable ", key)
	}
	return false
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
	return fmt.Sprintf("%s://%s:%d/%s?sslmode=disable&user=%s&password=%s",
		c.DriverName,
		c.Host,
		c.Port,
		c.Name,
		c.User,
		c.Password,
	)
}

func GetConfig() *Config {

	dbConfig := &DBConfig{
		Name:       getStrEnv("DB_NAME", "", true),
		DriverName: getStrEnv("DRIVER_NAME", "", true),
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
