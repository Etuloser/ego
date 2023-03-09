package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type (
	Config struct {
		MysqlSetting Mysql `mapstructure:"mysql"`
	}
	Mysql struct {
		Dsn string `mapstructure:"dsn"`
	}
)

// BaseDir -.
// 获取绝对路径
func BaseDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

// NewConfig -.
func NewConfig() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(BaseDir())
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
