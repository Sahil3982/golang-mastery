package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string
}
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flag := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flag

		if configPath == "" {
			panic("config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: ", err)
	}

	var cfg Config

	err := cleanenv.ReadCofig(configPath, &cfg)

	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}

	return &cfg

}
