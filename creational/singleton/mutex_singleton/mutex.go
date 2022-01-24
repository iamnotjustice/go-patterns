package mutex_singleton

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type config_MutexSingleton struct {
	addressableBytes int
}

var (
	mu       = &sync.Mutex{}
	instance *config_MutexSingleton
)

func GetAddressableBytes() int {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()

		// second nil-check in case some goroutine created an instance while we were acquiring the lock
		if instance == nil {
			viper.SetConfigName("cfg")
			viper.SetConfigType("json")
			viper.AddConfigPath(".")
			err := viper.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("fatal error config file: %w", err))
			}

			bytes := viper.GetInt("addressable_bytes")

			instance = &config_MutexSingleton{
				addressableBytes: bytes,
			}
		}
	}

	return instance.addressableBytes
}
