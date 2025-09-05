package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		Database `mapstructure:"database"`
	}

	Database struct {
		HostDatabase string `mapstructure:"host_database"`
		PortDatabase string `mapstructure:"port_database"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		DbName       string `mapstructure:"db_name"`
	}
)

func New() (*Config, error) {
	cfg := Config{}
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./internal/configs")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	for _, k := range v.AllKeys() {
		anyValue := v.Get(k)
		str, ok := anyValue.(string)
		if !ok {
			continue
		}

		replaced := os.ExpandEnv(str)
		v.Set(k, replaced)
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling file: %w", err))
	}

	return &cfg, nil
}
