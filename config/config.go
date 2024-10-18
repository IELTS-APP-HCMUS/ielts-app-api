package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("C:/Users/Admin/Desktop/project/ielts-app-api/config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config Config

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
