package mysql

import (
	"fmt"
	"getaway/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MysqlDB *gorm.DB
)

func Init(gCfg *config.AppConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		gCfg.Mysql.Username, gCfg.Mysql.Password, gCfg.Mysql.Path, gCfg.Mysql.Dbname)
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
