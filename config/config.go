package config

import (
	"github.com/redpkg/formula/db"
	"github.com/redpkg/formula/log"
	"github.com/redpkg/formula/redis"
	"github.com/spf13/viper"
	"strings"
)

type App struct {
	Addr         string `mapstructure:"addr"`
	Mode         string `mapstructure:"mode"`
	ReadTimeout  string `mapstructure:"read_timeout"`
	WriteTimeout string `mapstructure:"write_timeout"`
}

type Api struct {
	Rapidash string `mapstructure:"rapidash"`
	Ponyta   string `mapstructure:"ponyta"`
	Lugia    string `mapstructure:"lugia"`
}

type Conf struct {
	Project string       `mapstructure:"project"`
	App     App          `mapstructure:"app"`
	Log     log.Config   `mapstructure:"log"`
	DB      db.Config    `mapstructure:"db"`
	Redis   redis.Config `mapstructure:"redis"`
	Api     Api          `mapstructure:"api"`
}

func Init(files ...string) (conf *Conf, err error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, file := range files {
		viper.SetConfigFile(file)
		if err := viper.MergeInConfig(); err != nil {
			return nil, err
		}
	}

	conf = &Conf{}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
