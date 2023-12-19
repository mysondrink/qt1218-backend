/*
 * @Author: mysondrink
 * @Date: 2023-12-19 11:30:17
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-19 15:23:22
 * @Description:  主程序入口
 */

package main

import (
	"io"
	"os"
	"qt1218-backend/common"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDB()

	// 禁用控制台颜色， 将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()

	logFile := "gin.log"
	var f *os.File
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		f, _ = os.Create(logFile)
	} else {
		f, _ = os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0666)
	}

	defer f.Close()
	// 创建Multiwriter，包括文件和字符串缓冲区
	// 如果需要同时将日志写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

// 读取配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
