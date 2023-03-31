package config

import (
	"github.com/gestgo/gest/packages/common/config"
	"log"
)

var configuration *Configuration

type Configuration struct {
	Http config.HostPort
}

func init() {
	var err error
	configuration, err = config.LoadConfigYaml[Configuration]("./config/config.yaml")
	log.Print(configuration)
	if err != nil {
		log.Fatal(err)
	}
}
func GetConfiguration() *Configuration {
	return configuration
}
