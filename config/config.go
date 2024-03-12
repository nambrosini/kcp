package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	BootstrapServer  string `env:"BOOTSTRAP_SERVER,required"`
	KafkaCertificate string `env:"KAFKA_CERTIFICATE,required"`
	GroupId          string `env:"GROUP_ID"`
	Username         string `env:"USERNAME,required"`
	Password         string `env:"PASSWORD,required"`
}


func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load config: ", err)
	}

	cfg := Config{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	return cfg
}
