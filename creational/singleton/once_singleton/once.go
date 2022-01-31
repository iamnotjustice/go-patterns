package once_singleton

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type config_OnceSingleton struct {
	order string
}

var (
	once sync.Once

	instance *config_OnceSingleton
)

func GetOrder() string {
	once.Do(func() {
		viper.SetConfigName("cfg")
		viper.SetConfigType("json")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		bo := viper.GetString("byte_order")

		instance = &config_OnceSingleton{
			order: bo,
		}
	})

	return instance.order
}
