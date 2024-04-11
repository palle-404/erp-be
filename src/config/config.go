package config

import "github.com/spf13/viper"

var appCfg *viper.Viper

func AppCfg() *viper.Viper {
	return appCfg
}

func Load() (err error) {
	tmpCfg := viper.New()
	tmpCfg.AddConfigPath(".")
	tmpCfg.SetConfigFile("config.yaml")
	if err := tmpCfg.ReadInConfig(); err != nil {
		return err
	}
	appCfg = tmpCfg
	return nil
}
