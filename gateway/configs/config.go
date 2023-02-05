package configs

import (
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	ErrorCode []struct {
		ID      int    `mapstructure:"ID"`
		Message string `mapstructure:"MESSAGE"`
	}
}

var (
	conf Config
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {
	runtime_viper := viper.New()
	runtime_viper.AddRemoteProvider("consul", "consul-headless.consul.svc.cluster.local:8500", "gateway")

	runtime_viper.SetConfigType("json") // Need to explicitly set this to json
	runtime_viper.ReadRemoteConfig()

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				continue
			}
			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			runtime_viper.Unmarshal(&conf)
		}
	}()

	return &conf
}
