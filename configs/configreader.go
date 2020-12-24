package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

const file string = "config.yml"
const path string = ".."
const configType string = "yml"

func Read() (Configurations, error) {

	viper.SetConfigFile(file)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigType(configType)

	var config Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)

	return config, err
}
