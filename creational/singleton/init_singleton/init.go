package init_singleton

import (
	"fmt"

	"github.com/spf13/viper"
)

type config_InitSingleton struct {
	arch string
}

var (
	instance *config_InitSingleton
)

func init() {
	instance = new()
}

func new() *config_InitSingleton {
	viper.SetConfigName("cfg")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	arch := viper.GetString("arch")

	return &config_InitSingleton{
		arch: arch,
	}
}

func GetArch() string {
	return instance.arch
}
