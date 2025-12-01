package mysql

import (
	"database/sql"
	"fmt"
	"web_app_go/settings"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db    *gorm.DB
	sqlDB *sql.DB
)

func InitMysql(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		zap.L().Error("init mysql fail", zap.Error(err))
	}
	sqlDB, err = Db.DB()
	if err != nil {
		zap.L().Error("init mysql fail", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifeTime)
	return
}

// Close 关闭数据库连接
func Close() {
	_ = sqlDB.Close()
}
