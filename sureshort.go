package main

import (
	"log"
	"path"

	"github.com/spf13/viper"
	"github.com/surajbhosale409/sureshort/service"
)

var configDir = "config"

func loadConfigFromFile() (*service.Config, error) {
	viper.SetConfigName(path.Join(configDir, "config.defaults"))
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg service.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	config, err := loadConfigFromFile()
	if err != nil {
		log.Fatal(err)
	}

	urlShortnerService := service.NewService(config)
	urlShortnerService.Serve()
}
