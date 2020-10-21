package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// C ...
var C *Schema

// New ...
func New(configPath string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/m9")
	if len(configPath) != 0 {
		viper.AddConfigPath(configPath)
	}
	viper.AddConfigPath(".")

	errReadInConfig := viper.ReadInConfig()
	if errReadInConfig != nil {
		panic(errReadInConfig)
	}

	var configSchema Schema
	errUnmarshal := viper.Unmarshal(&configSchema)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	C = &configSchema

	fmt.Printf("%+v\n", C)
}
