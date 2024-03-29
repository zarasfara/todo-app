package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	defaultHttpPort = "8000"
)

type (
	Config struct {
		DBConnection DBConfig
		HTTP         HTTPConfig
	}

	DBConfig struct {
		Host     string
		Port     string
		Database string
		Username string
		Password string
		SSLMode  string
	}

	HTTPConfig struct {
		Port string
	}
)

// Init confgire from .env and config.yml into Config struct.
func Init() *Config {
	populateDefaults()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	if err := parseConfig(); err != nil {
		log.Fatal("Cannot parse config")
	}

	var cfg Config

	setFromEnv(&cfg)

	if err := unmarshal(&cfg); err != nil {
		log.Fatalf("Cannot unmarshal config: %s", err.Error())
	}

	return &cfg
}

// unmarshal set parametres from config.yml file.
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		log.Fatalf("Error read config: %s", err)
	}
	if err := viper.UnmarshalKey("db", &cfg.DBConnection); err != nil {
		log.Fatalf("Error read config: %s", err)
	}

	return nil
}

// setFromEnv set parametres from .env.
func setFromEnv(cfg *Config) {
	// database
	cfg.DBConnection.Database = os.Getenv("DB_DATABASE")
	cfg.DBConnection.Password = os.Getenv("DB_PASSWORD")
	cfg.DBConnection.Username = os.Getenv("DB_USERNAME")
	cfg.DBConnection.Port = os.Getenv("DB_PORT")
	cfg.DBConnection.Host = os.Getenv("DB_HOST")
}

// populateDefaults set default parametres for config.
func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
}

// parseConfig parse config from file config.yml.
func parseConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
