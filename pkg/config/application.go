package config

import "github.com/spf13/viper"

type Application struct {
	Mode string `json:"mode"`
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"name"`
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		Mode: cfg.GetString("mode"),
		Host: cfg.GetString("host"),
		Port: cfg.GetString("port"),
		Name: cfg.GetString("name"),
	}
}

var ApplicationConfig = new(Application)
