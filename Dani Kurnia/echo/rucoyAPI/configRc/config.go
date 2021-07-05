package configRc

import (
	"github.com/tkanos/gonfig"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Config {
	conf := Config{}
	gonfig.GetConf("configRc/config.json", &conf)
	return conf
}
