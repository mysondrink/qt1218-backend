/*
 * @Author: mysondrink
 * @Date: 2023-12-19 14:35:48
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-19 17:24:44
 * @Description:  数据库配置连接处理
 */

package common

import (
	"fmt"
	"qt1218-backend/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 使用viper读取数据库配置信息
	// driverName := "mysql"
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username, password, host, port, database, charset, loc)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// table name prefix, table for `User` would be `t_users`
			// TablePrefix:   "t_",
			// use singular table name, table for `User` would be `user` with this option enabled
			SingularTable: true,
			// skip the snake_casing of names
			NoLowerCase: false,
		},
	})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	db.AutoMigrate(&model.UserTable{}, &model.ReagentCopy1{}, &model.PatientCopy1{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
