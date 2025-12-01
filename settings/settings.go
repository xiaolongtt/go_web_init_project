package settings

//用来管理配置的文件，使用viper进行管理配置

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig() (err error) {
	viper.SetConfigFile("./config.yaml")
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig() //实现热更新热加载
	//更新配置后的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	return
}
