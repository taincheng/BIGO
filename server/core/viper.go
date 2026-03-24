package core

import (
	"flag"
	"fmt"
	"os"
	"server/core/internal"
	"server/global"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	configPath := getConfigPath()
	v := viper.New()
	v.SetConfigName(configPath)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err = v.Unmarshal(&global.BIGO_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.BIGO_CONFIG); err != nil {
		panic(fmt.Errorf("unmarshal config err: %w", err))
	}

	return v
}

// getConfigPath 获取配置文件路径, 优先级 命令行参数 > 环境变量 > 默认值
func getConfigPath() (config string) {
	// -c 指定配置文件路径
	flag.StringVar(&config, "c", "", "choose config file")
	if config != "" {
		fmt.Printf("use command line -c to get config path: %s\n", config)
		return
	}

	// 获取环境变量
	if env := os.Getenv(internal.ConfigEnv); env != "" {
		config = env
		fmt.Printf("use environment variable %s to get config path: %s\n", internal.ConfigEnv, config)
		return
	}

	// 根据设置的环境变量：GIN_MODE，来获取默认的配置文件
	switch gin.Mode() {
	case gin.DebugMode:
		config = internal.ConfigDebugFile
	case gin.ReleaseMode:
		config = internal.ConfigReleaseFile
	case gin.TestMode:
		config = internal.ConfigTestFile
	}
	fmt.Printf("use gin mode: %s, config path is: %s\n", gin.Mode(), config)
	if _, err := os.Stat(config); err != nil || os.IsNotExist(err) {
		config = internal.ConfigDefaultFile
		fmt.Printf("use default config path: %s\n", config)
	}
	return
}
