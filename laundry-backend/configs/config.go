package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Log      LogConfig
}

type ServerConfig struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret string
	Expire int
}

type LogConfig struct {
	Level string
}

func LoadConfig() (*Config, error) {

	exp, _ := strconv.Atoi(GetEnv("JWT_EXPIRE"))
	readTimeout, _ := strconv.Atoi(GetEnv("SERVER_READ_TIMEOUT"))
	writeTimeout, _ := strconv.Atoi(GetEnv("SERVER_WRITE_TIMEOUT"))
	config := &Config{
		Server: ServerConfig{
			Address:      GetEnv("APP_PORT"),
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
		Database: DatabaseConfig{
			Host:     GetEnv("DB_HOST"),
			Port:     GetEnv("DB_PORT"),
			User:     GetEnv("DB_USER"),
			Password: GetEnv("DB_PASSWORD"),
			Name:     GetEnv("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret: GetEnv("JWT_SECRET"),
			Expire: exp,
		},
		Log: LogConfig{
			Level: GetEnv("LOG_LEVEL"),
		},
	}

	// Debug: Print individual config values
	log.Printf("DB Config: host=%s port=%s user=%s password=%s name=%s",
		config.Database.Host, config.Database.Port, config.Database.User,
		config.Database.Password, config.Database.Name)

	// Validasi konfigurasi
	if config.Server.Address == "" {
		return nil, fmt.Errorf("SERVER_ADDRESS tidak boleh kosong")
	}

	return config, nil
}

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Load file .env not found")
	}

	if os.Getenv(key) != "" {
		log.Println(key, os.Getenv(key))
		return os.Getenv(key)
	} else {
		if len(value) > 0 {
			log.Println(key, value)
			return value[0]
		}
		return ""
	}
}
