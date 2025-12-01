package mysql

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db    *gorm.DB
	sqlDB *sql.DB
)

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("datasource.mysql.username"),
		viper.GetString("datasource.mysql.password"),
		viper.GetString("datasource.mysql.host"),
		viper.GetString("datasource.mysql.port"),
		viper.GetString("datasource.mysql.dbName"),
	)
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		zap.L().Error("init mysql fail", zap.Error(err))
	}
	sqlDB, err = Db.DB()
	if err != nil {
		zap.L().Error("init mysql fail", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(viper.GetInt("datasource.mysql.max_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("datasource.mysql.max_open_conns"))
	sqlDB.SetConnMaxLifetime(viper.GetDuration("datasource.mysql.max_lifetime"))
	return
}

// Close 关闭数据库连接
func Close() {
	_ = sqlDB.Close()
}
