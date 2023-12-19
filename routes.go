/*
 * @Author: mysondrink
 * @Date: 2023-12-19 14:36:36
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-19 15:41:38
 * @Description:  路由处理
 */
package main

import (
	"qt1218-backend/middleware"

	"github.com/gin-gonic/gin"
)

// 定义路由
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/api/reagent/info", controller.GetReagentInfo)

	return r
}
