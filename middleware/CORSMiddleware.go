/*
 * @Author: mysondrink
 * @Date: 2023-12-19 14:44:23
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-19 15:39:28
 * @Description:  浏览器跨域请求问题
 */
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置域名信息，包括可以访问的域名信息，“*”为全部访问域名
		ctx.Writer.Header().Set("Access-control-Allow-Origin", "http://localhost:8080")
		// 设置缓存时间
		ctx.Writer.Header().Set("Access-control-Max-Age-Origin", "86400")
		// 设置可以通过访问的方法
		// GET，POST
		ctx.Writer.Header().Set("Access-control-Allow-Methods", "*")
		// 设置可以通过访问的头
		ctx.Writer.Header().Set("Access-control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
