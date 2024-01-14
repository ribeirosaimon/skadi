package skadiEngine

import (
	"os"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/skadi/domain/config"
)

var (
	p   *properties.Properties
	env string
)

func GetProperties() *properties.Properties {
	env = os.Args[1]
	p = config.GetPropertiesFile(env)
	return p
}

func GetEnvironment() string {
	return env
}
