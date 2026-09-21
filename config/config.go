package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type DBConfig struct{
	Host string `yml:"host"`
	Port string `yml:"port"`
	Username string 
	Password string 
	DBName string `yml:"dbname"`
	SSlMode string `yml:"sslmode"`
}

type Config struct{
	Port string `yml:"port" env:"PORT" env-default:"6767"`
	YTTOKEN string `env:"YOUTUBETOKEN" `
	DB DBConfig `yml:"db"`
}

func YMLConfig(conf *Config) error {
	return cleanenv.ReadConfig("./config/config.yml", conf)

}
func EnvConfig(conf *Config) error {
	if err := godotenv.Load(".env"); err != nil{
		return err
	}
	conf.YTTOKEN = os.Getenv("YTTOKEN")
	conf.DB.Password = os.Getenv("DBPASSWORD")
	conf.DB.Username = os.Getenv("DBUSERNAME")
	return nil
}

func NewConfig() (*Config, error){
	var conf Config

	if err := YMLConfig(&conf); err != nil{
		return nil, err
	}

	if err := EnvConfig(&conf); err != nil{
		return nil, err
	}
	log.Println("DKJFSDJFJKSDF", conf.DB.Port)
	return &conf, nil
}
