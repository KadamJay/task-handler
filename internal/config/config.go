package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type MongoDBConfig struct {
	URI string `yaml:"uri"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type Config struct {
	MongoDB MongoDBConfig `yaml:"mongodb"`
	Server  ServerConfig  `yaml:"server"`
}

func LoadConfig(configPath string) *Config {
	file, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error Reading config file, %s", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

	overriderWithEnvVariable(&config)
	return &config
}

func overriderWithEnvVariable(config *Config) {
	if uri := os.Getenv("MONGODB_URI"); uri != "" {
		config.MongoDB.URI = uri
	}

	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}
}
