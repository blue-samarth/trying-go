package config 

import (
	"fmt"
	"os"
	"log"
	"flag"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"dev" `
	Storagepath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true" env-default:"/storage/storage.db"`
	HTTPServer HTTPServer `yaml:"http_server" env:"HTTP_SERVER" env-required:"true" `
}

func MustLoad() *Config {
	// Load the configuration from a file or environment variables
	
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if config.Env == "" {
		flag := flag.String("env", "dev", "Environment (dev, test, prod)")
		flag.Parse()

		config.Env = *flag

		if config.Env == "" {
			log.Fatal("ENV is required")
		}
	}
	if configPath == "" {
		flag := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flag
		if configPath == "" {
			log.Fatal("CONFIG_PATH is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}


	if config.HTTPServer.Address == "" {
		panic("HTTP_SERVER_ADDRESS is required")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {	log.Fatalf("Error loading config file: %v", err.Error()) }

	fmt.Println("Configuration loaded successfully:", config)

	return &cfg
}