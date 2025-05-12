package config 

import (
	"fmt"
	"os"
	"log"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"`
	Storagepath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true" env-default:"/storage/storage.db"`
	HTTPServer HTTPServer `yaml:"http_server" env:"HTTP_SERVER" env-required:"true"`
}

func MustLoad() *Config {
	// Load the configuration from a file or environment variables
	
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configflag := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *configflag
		if configPath == "" {
			log.Fatal("CONFIG_PATH is required")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}


	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {	log.Fatalf("Error loading config file: %v", err.Error()) }

	fmt.Println("Configuration loaded successfully")

	return &cfg
}