/*
 * @Author: mysondrink
 * @Date: 2023-12-19 16:08:08
 * @Last Modified by: mysondrink
 * @Last Modified time: 2023-12-19 17:29:11
 * @Description:  发送响应信息
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
