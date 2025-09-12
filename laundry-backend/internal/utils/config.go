package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
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
	// viper.SetConfigFile("./.env")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("/app") // explicit
	viper.AddConfigPath(".")    // current dir
	viper.AddConfigPath("..")   // parent dir

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	// viper.AutomaticEnv() // Removed to prevent environment variables from overriding .env

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Debug: Print all config values
	log.Printf("All config values: %+v", viper.AllSettings())

	config := &Config{
		Server: ServerConfig{
			Address:      viper.GetString("SERVER_ADDRESS"),
			ReadTimeout:  viper.GetInt("SERVER_READ_TIMEOUT"),
			WriteTimeout: viper.GetInt("SERVER_WRITE_TIMEOUT"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Name:     viper.GetString("DB_NAME"),
		},
		JWT: JWTConfig{
			Secret: viper.GetString("JWT_SECRET"),
			Expire: viper.GetInt("JWT_EXPIRE"),
		},
		Log: LogConfig{
			Level: viper.GetString("LOG_LEVEL"),
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
