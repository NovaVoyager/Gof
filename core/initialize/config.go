package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/miaogu-go/Gof/global"
	"github.com/spf13/viper"
	"time"
)

//LoadConfig 加载配置文件
func LoadConfig() {
	viper.SetConfigFile(global.ConfigFile)
	viper.SetConfigType(global.ConfigType)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}

	//读取到的配置文件存入结构体
	err = viper.Unmarshal(&global.GofConfig)
	if err != nil {
		panic(fmt.Errorf("Unmarshal config error:%s \n", err))
	}

	//监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config edit", time.Now().Format("2006-01-02 15:04:05"))
		err = viper.Unmarshal(&global.GofConfig)
		if err != nil {
			panic(fmt.Errorf("Unmarshal config error:%s \n", err))
		}
		LoadLog()
		LoadDb()
		LoadRedis()
		LoadMongodb()
	})
}
