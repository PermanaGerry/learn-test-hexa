package config

import "github.com/spf13/viper"

func LoadEnvVar() {
	viper.SetConfigFile("../../config/.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.SetDefault("TIMEZONE", "UTC")
}
