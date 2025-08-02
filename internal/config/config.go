package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string
}

// env-default:"production"
// Stage tags 

type Config struct {
    Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var ConfigPath string 
	
	ConfigPath = os.Getenv("CONFIG_PATH")

	if ConfigPath == "" {
		flags := flag.String("config" , "" , "path to the configuration file") 
		flag.Parse()

		ConfigPath= *flags

		if ConfigPath == "" {
			log.Fatal("Config Path Is Not Set")
		}
	}

	if _,err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatalf("Config Path Is Not Set %s", ConfigPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(ConfigPath , &cfg)

	if err != nil {
		log.Fatalf("Can not read config file: %s" , err.Error())
	}

	return  &cfg
}