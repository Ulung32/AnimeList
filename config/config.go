package config

import "github.com/spf13/viper"

type Config struct {
	DBUrl     string `mapstructure:"DATABASE_URL"`
	SecretKey string `mapstructure:"SECRET"`
}

var Cfg *Config

func LoadConfig() (err error) {
	viper.AddConfigPath("./.")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	Cfg = &Config{}

	err = viper.Unmarshal(Cfg)
	return
}
