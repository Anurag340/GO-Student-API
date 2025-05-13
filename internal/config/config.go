package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTP struct {
	Addr string `yaml:"address" env-required:"true"`
}

// struct-tags
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTP        `yaml:"http_server"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags:= flag.String("config", "", "Path to the config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("config path is required")
		}
	}

	if _,err :=os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file does not exist: %s" , configPath)
	}

	var cfg Config

	err:=cleanenv.ReadConfig(configPath, &cfg)

	if err!=nil{
		log.Fatalf("can not read config file: %s",err.Error())

	}

	return &cfg
}