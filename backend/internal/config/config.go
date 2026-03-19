package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

var Cfg = &Config{}

func Init() error {
	// 配置文件路径配置
	viper.AddConfigPath("./backend/config")
	// 获取环境变量标识，读取不同的配置文件
	env := os.Getenv("BIGO_ENV")
	if env == "" {
		env = "dev"
	}
	configPath := fmt.Sprintf("config.%s", env)
	viper.SetConfigName(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %s err: %w", configPath, err)
	}
	err = viper.Unmarshal(Cfg)
	if err != nil {
		return fmt.Errorf("unmarshal config err: %w", err)
	}

	// 监控配置文件，实时更新
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		err = viper.Unmarshal(Cfg)
	})
	viper.WatchConfig()

	return nil
}
