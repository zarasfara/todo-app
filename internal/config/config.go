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
		Database string
		Username string
		Password string
	}

	HTTPConfig struct {
		Port string
	}
)

// init configuration from .env and config.yml into Config struct.
func Init() *Config {
	populateDefaults()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := parseConfig(); err != nil {
		log.Fatal("Cannot parse config")
	}

	var cfg Config

	setFromEnv(&cfg)

	if err := unmarshal(&cfg); err != nil {
		log.Fatal("Cannot unmarshal config")
	}


	return &cfg
}

// set parametres from config.yml file.
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http.port", cfg.HTTP); err != nil {
		log.Fatal("Error read config")
	}

	return nil
}

// set parametres from .env.
func setFromEnv(cfg *Config) {
	// database
	cfg.DBConnection.Database = os.Getenv("db_pet")
	cfg.DBConnection.Password = os.Getenv("root")
	cfg.DBConnection.Username = os.Getenv("root")
}

// set default parametres for config.
func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
}

// parse config from file config.yml.
func parseConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
