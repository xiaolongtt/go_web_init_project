package main

import (
	"fmt"
	"web_app_go/dao/mysql"
	"web_app_go/dao/redis"
	"web_app_go/logger"
	"web_app_go/route"
	"web_app_go/settings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// go_web开发通用脚手架
func main() {
	//1.加载配置文件
	if err := settings.LoadConfig(); err != nil {
		fmt.Printf("init config file failed,err:%v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.InitLogger(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	zap.L().Debug("init logger success")
	//3.连接数据库，redis和mysql
	if err := mysql.InitMysql(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	if err := redis.InitRedis(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return
	}
	//4.注册路由
	r := route.InitRouter()
	//5.启动服务
	err := r.Run(":" + viper.GetString("server.port"))
	if err != nil {
		return
	}
}
