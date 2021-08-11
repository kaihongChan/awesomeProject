package initialize

import (
	"awesomeProject/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ViperInit() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("config file not found: %w \n", err))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}

	v.WatchConfig()
	// 监听配置文件修改
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
