package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
