package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/etc"
)

func InitMysql(handle func(db *gorm.DB) error) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		etc.Conf.Mysql.User,
		etc.Conf.Mysql.Pass,
		etc.Conf.Mysql.Hort,
		etc.Conf.Mysql.Port,
		etc.Conf.Mysql.Dbname,
	)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db, err := cli.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	return handle(cli)
}
